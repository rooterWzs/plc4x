/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.omronfins.optimizer;

import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.omronfins.context.OmronFinsDriverContext;
import org.apache.plc4x.java.omronfins.readwrite.DataItem;
import org.apache.plc4x.java.omronfins.readwrite.OmronFinsDataType;
import org.apache.plc4x.java.omronfins.tag.OmronFinsTag;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadRequest;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadResponse;
import org.apache.plc4x.java.spi.messages.PlcReader;
import org.apache.plc4x.java.spi.messages.utils.DefaultPlcResponseItem;
import org.apache.plc4x.java.spi.messages.utils.DefaultPlcTagItem;
import org.apache.plc4x.java.spi.messages.utils.PlcResponseItem;
import org.apache.plc4x.java.spi.messages.utils.PlcTagItem;
import org.apache.plc4x.java.spi.optimizer.SingleTagOptimizer;
import org.apache.plc4x.java.spi.values.PlcBOOL;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.*;

/**
 * In order to read more data more efficiently, this optimizer for modbus joins together individual items
 * and reads larger arrays of data.
 */
public class FinsOptimizer extends SingleTagOptimizer {
    private final Logger logger = LoggerFactory.getLogger(FinsOptimizer.class);

    /**
     * Per default the number of registers that can be read are 125 registers.
     * The number of coils that can be read in one request are 2000 coils.
     * This optimizer will sort all request items, split them up into coils and registers.
     * Then it will try to group them together to produce the least amount of traffic
     * on the wire by requesting larger chunks of words and serving multiple items
     * with one request. The matching processReadResponses will then handle the splitting
     * of the results.
     *
     * @param readRequest   the original read request
     * @param driverContext the driver context
     * @return a list of rewritten sub-requests
     */
    @Override
    protected List<PlcReadRequest> processReadRequest(PlcReadRequest readRequest, DriverContext driverContext) {
        OmronFinsDriverContext modbusContext = (OmronFinsDriverContext) driverContext;

        // Sort the different types of tags as all need to be requested separately anyway.
        TreeSet<OmronFinsTag> coils = null;
        TreeSet<OmronFinsTag> holdingRegisters = null;

        for (PlcTag tag : readRequest.getTags()) {
            if (tag instanceof OmronFinsTag) {
                if (coils == null) {
                    coils = new TreeSet<>(Comparator.comparingInt(OmronFinsTag::getAddrword));
                }
                coils.add(((OmronFinsTag) tag));
            } else if (tag instanceof OmronFinsTag) {
                if (holdingRegisters == null) {
                    holdingRegisters = new TreeSet<>(Comparator.comparingInt(OmronFinsTag::getAddrword));
                }
                holdingRegisters.add(((OmronFinsTag) tag));
            } 
        }

        // Add sub requests for every type of tag in this request.
        PlcReader reader = ((DefaultPlcReadRequest) readRequest).getReader();
        List<PlcReadRequest> subRequests = new ArrayList<>();
        if (coils != null) {
            subRequests.addAll(processBitRequests(coils, reader, modbusContext));
        }
//        if (holdingRegisters != null) {
//            subRequests.addAll(processByteRequests(holdingRegisters, reader, (address, count, dataType) -> new OmronFinsTagHoldingRegister(address, count, dataType, Collections.emptyMap()), modbusContext));
//        }
        
        return subRequests;
    }

    /**
     * When reading large chunks of data, here we need to read the parts that were originally requested.
     *
     * @param readRequest original read request
     * @param readResponses map of the sub-requests that were executed
     * @return post processed response reflecting the unmodified user request items
     */
    @Override
    protected PlcReadResponse processReadResponses(PlcReadRequest readRequest, Map<PlcReadRequest, SubResponse<PlcReadResponse>> readResponses, DriverContext driverContext) {
        OmronFinsDriverContext modbusContext = (OmronFinsDriverContext) driverContext;
        try {
            // Build an index of all the data returned by all requests.
            // This data should contain all the bits needed to create the response of the original request.
            Map<String, List<Response>> responses = new HashMap<>();
            for (PlcReadRequest optimizedReadRequest : readResponses.keySet()) {
                PlcReadResponse optimizedReadResponse = readResponses.get(optimizedReadRequest).getResponse();
                if (optimizedReadResponse == null) {
                    continue;
                }
                // Optimized read requests only contain one OmronFinsTag.
                String tagName = optimizedReadRequest.getTagNames().stream().findFirst().orElse(null);
                if (tagName == null) {
                    continue;
                }
                OmronFinsTag OmronFinsTag = (OmronFinsTag) optimizedReadRequest.getTag(tagName);
                String tagType = OmronFinsTag.getClass().getSimpleName().substring("OmronFinsTag".length());
                if (!responses.containsKey(tagType)) {
                    responses.put(tagType, new ArrayList<>());
                }
                PlcResponseCode responseCode = optimizedReadResponse.getResponseCode(tagName);
                int startingAddress = OmronFinsTag.getAddrword();
                int endingAddressRegister = OmronFinsTag.getAddrword() + OmronFinsTag.getNumberOfElements();
                int endingAddressCoil = OmronFinsTag.getAddrword() + OmronFinsTag.getNumberOfElements();
                byte[] responseData = (responseCode == PlcResponseCode.OK) ? optimizedReadResponse.getPlcValue(tagName).getRaw() : null;
                responses.get(tagType).add(new Response(responseCode, startingAddress,
                    endingAddressRegister, endingAddressCoil, responseData));
            }

            // Now go through the original requests and try to answer them by using the raw data we now have.
            Map<String, PlcResponseItem<PlcValue>> values = new HashMap<>();
            for (String tagName : readRequest.getTagNames()) {
                OmronFinsTag OmronFinsTag = (OmronFinsTag) readRequest.getTag(tagName);
                String tagType = OmronFinsTag.getClass().getSimpleName().substring("OmronFinsTag".length());
                if (!responses.containsKey(tagType)) {
                    values.put(tagName, new DefaultPlcResponseItem<>(PlcResponseCode.NOT_FOUND, null));
                    continue;
                }
                // Go through all responses till we find one where that contains the current tag's data.
                for (Response response : responses.get(tagType)) {
                    if(OmronFinsTag instanceof OmronFinsTag) {
                        if(response.matchesCoil(OmronFinsTag)) {
                            // If this response was invalid, return all associated addresses as equally invalid.
                            // TODO: Possibly it would be worth doing a single item request for each of these
                            //  tags in order to find out which ones are actually invalid as if one item in the
                            //  current request exceeds the address range, all items in this chunk will fail, even
                            //  if only one element was invalid.
                            if(response.getResponseCode() != PlcResponseCode.OK) {
                                values.put(tagName, new DefaultPlcResponseItem<>(response.getResponseCode(), null));
                                break;
                            }

                            // Coils are read completely different from registers.
                            OmronFinsTag coilTag = (OmronFinsTag) OmronFinsTag;

                            // Calculate the byte that contains the response for this Coil
                            byte[] responseData = response.getResponseData();
                            int bitPosition = coilTag.getAddrword() - response.startingAddress;
                            int bytePosition = bitPosition / 8;
                            int bitPositionInByte = bitPosition % 8;
                            boolean isBitSet = (responseData[bytePosition] & (1 << bitPositionInByte)) != 0;
                            values.put(tagName, new DefaultPlcResponseItem<>(PlcResponseCode.OK, new PlcBOOL(isBitSet)));
                            break;
                        }
                    }
                    // Read a normal register.
                    else if (response.matchesRegister(OmronFinsTag)) {
                        // If this response was invalid, return all associated addresses as equally invalid.
                        // TODO: Possibly it would be worth doing a single item request for each of these
                        //  tags in order to find out which ones are actually invalid as if one item in the
                        //  current request exceeds the address range, all items in this chunk will fail, even
                        //  if only one element was invalid.
                        if(response.getResponseCode() != PlcResponseCode.OK) {
                            values.put(tagName, new DefaultPlcResponseItem<>(response.getResponseCode(), null));
                            break;
                        }

//                        byte[] responseData = response.getResponseDataForTag(OmronFinsTag);
//                        ReadBuffer readBuffer = getReadBuffer(responseData, modbusContext.getByteOrder());
//                        try {
//                            PlcValue plcValue = DataItem.staticParse(readBuffer, OmronFinsTag.getDataType(),
//                                OmronFinsTag.getNumberOfElements(),
//                                modbusContext.getByteOrder() == ModbusByteOrder.BIG_ENDIAN);
//                            values.put(tagName, new DefaultPlcResponseItem<>(PlcResponseCode.OK, plcValue));
//                        } catch (ParseException e) {
//                            values.put(tagName, new DefaultPlcResponseItem<>(PlcResponseCode.INTERNAL_ERROR, null));
//                        }
                        break;
                    }
                }
                // If no response was found that contains the data, that's probably something we need to fix.
                if (!values.containsKey(tagName)) {
                    values.put(tagName, new DefaultPlcResponseItem<>(PlcResponseCode.INTERNAL_ERROR, null));
                }
            }

            return new DefaultPlcReadResponse(readRequest, values);
        }
        // TODO: Remove this part once the driver is more stable.
        catch (Exception e) {
            try {
                logger.error("Error processing response:", e);
                WriteBufferXmlBased wb = new WriteBufferXmlBased();
                ((DefaultPlcReadRequest) readRequest).serialize(wb);
                logger.error("Original Request:\n{}", wb.getXmlString(), e);
                for (PlcReadRequest subRequest : readResponses.keySet()) {
                    SubResponse<PlcReadResponse> subResponse = readResponses.get(subRequest);
                    WriteBufferXmlBased wbSubRequest = new WriteBufferXmlBased();
                    ((DefaultPlcReadRequest) subRequest).serialize(wbSubRequest);
                    logger.error("Sub Request:\n{}", wbSubRequest.getXmlString());
                    if(subResponse.isSuccess()) {
                        PlcReadResponse plcReadResponse = subResponse.getResponse();
                        WriteBufferXmlBased wbSubResponse = new WriteBufferXmlBased();
                        ((DefaultPlcReadResponse) plcReadResponse).serialize(wbSubResponse);
                        logger.error("Sub Response (Success):\n{}", wbSubResponse.getXmlString());
                    } else {
                        Throwable throwable = subResponse.getThrowable();
                        logger.error("Sub Response (Error):", throwable);
                    }
                }
            } catch (SerializationException ex) {
                throw new RuntimeException(ex);
            }
            return null;
        }
    }

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Internal
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    protected List<PlcReadRequest> processBitRequests(TreeSet<OmronFinsTag> tags, PlcReader reader, OmronFinsDriverContext modbusContext) {
        List<PlcReadRequest> subRequests = new ArrayList<>();
        int firstCoil = -1;
        int lastCoil = -1;
        int maxCoilCurRequest = -1;
        for (OmronFinsTag tag : tags) {
            int sizeInCoils = tag.getDataType().getDataTypeSize() * 8;
            if (tag.getDataType() == OmronFinsDataType.BOOL) {
                sizeInCoils = 1;
            }
            // Initialize for the first item.
            if (firstCoil == -1) {
                firstCoil = tag.getAddrword();
                lastCoil = tag.getAddrword() + (sizeInCoils * tag.getNumberOfElements());
                // 2000 coils/request is the modbus limit.
//                maxCoilCurRequest = tag.getAddrword() + modbusContext.getMaxCoilsPerRequest();
            }

            // If adding the current coil would exceed the maximum number of coils that can be read by one request,
            // finish this one and start a new one.
            if (tag.getAddrword() + (sizeInCoils * tag.getNumberOfElements()) > maxCoilCurRequest) {
                // Finish the current sub-request
                LinkedHashMap<String, PlcTagItem<PlcTag>> subTags = new LinkedHashMap<>();
//                subTags.put("coils" + subRequests.size(), new DefaultPlcTagItem<>(new OmronFinsTag(firstCoil, lastCoil - firstCoil, ModbusDataType.BYTE, Collections.emptyMap())));
//                subRequests.add(new DefaultPlcReadRequest(reader, subTags));

                // Re-initialize the structures for the next request.
                firstCoil = tag.getAddrword();
                lastCoil = tag.getAddrword() + (sizeInCoils * tag.getNumberOfElements());
//                maxCoilCurRequest = tag.getAddrword() + modbusContext.getMaxCoilsPerRequest();
            }
            // Otherwise update the end-marker for the current block.
            else {
                lastCoil = tag.getAddrword() + tag.getNumberOfElements();
            }
        }

        // Finish the last sub-request
        LinkedHashMap<String, PlcTagItem<PlcTag>> subTags = new LinkedHashMap<>();
//        subTags.put("coils" + subRequests.size(), new DefaultPlcTagItem<>(new OmronFinsTag(firstCoil, lastCoil - firstCoil, ModbusDataType.BYTE, Collections.emptyMap())));
//        subRequests.add(new DefaultPlcReadRequest(reader, subTags));
        return subRequests;
    }

    protected List<PlcReadRequest> processByteRequests(TreeSet<OmronFinsTag> tags, PlcReader reader, TagFactory tagFactory, OmronFinsDriverContext modbusContext) {
        List<PlcReadRequest> subRequests = new ArrayList<>();
        int firstRegister = -1;
        int lastRegister = -1;
        int maxRegisterCurRequest = -1;
        for (OmronFinsTag tag : tags) {
            int sizeInRegisters = (int) Math.ceil((double) tag.getDataType().getDataTypeSize() / 2);
            // Initialize for the first item.
            if (firstRegister == -1) {
                firstRegister = tag.getAddrword();
                lastRegister = tag.getAddrword() + (sizeInRegisters * tag.getNumberOfElements());
                // 2000 coils/request is the modbus limit.
//                maxRegisterCurRequest = tag.getAddrword() + modbusContext.getMaxRegistersPerRequest();
            }

            // If adding the current coil would exceed the maximum number of coils that can be read by one request,
            // finish this one and start a new one.
            if (tag.getAddrword() + (sizeInRegisters * tag.getNumberOfElements()) > maxRegisterCurRequest) {
                // Finish the current sub-request
                LinkedHashMap<String, PlcTagItem<PlcTag>> subTags = new LinkedHashMap<>();
//                subTags.put("registers" + subRequests.size(), new DefaultPlcTagItem<>(tagFactory.createTag(firstRegister, lastRegister - firstRegister, ModbusDataType.WORD)));
//                subRequests.add(new DefaultPlcReadRequest(reader, subTags));

                // Re-initialize the structures for the next request.
                firstRegister = tag.getAddrword();
                lastRegister = tag.getAddrword() + (sizeInRegisters * tag.getNumberOfElements());
//                maxRegisterCurRequest = tag.getAddrword() + modbusContext.getMaxRegistersPerRequest();
            }
            // Otherwise update the end-marker for the current block.
            else {
                lastRegister = tag.getAddrword() + (sizeInRegisters * tag.getNumberOfElements());
            }
        }

        // Finish the last sub-request
        LinkedHashMap<String, PlcTagItem<PlcTag>> subTags = new LinkedHashMap<>();
//        subTags.put("registers" + subRequests.size(), new DefaultPlcTagItem<>(tagFactory.createTag(firstRegister, lastRegister - firstRegister, ModbusDataType.WORD)));
//        subRequests.add(new DefaultPlcReadRequest(reader, subTags));
        return subRequests;
    }

    protected static class Response {
        private final PlcResponseCode responseCode;
        private final int startingAddress;
        private final int endingAddressCoil;
        private final int endingAddressRegister;
        private final byte[] responseData;

        public Response(PlcResponseCode responseCode, int startingAddress, int endingAddressRegister, int endingAddressCoil, byte[] responseData) {
            this.responseCode = responseCode;
            this.startingAddress = startingAddress;
            this.responseData = responseData;
            this.endingAddressCoil = endingAddressRegister;
            // In general a "Celil(responseData.length / 2)" would have been more correct,
            // but the data returned from the device should already be an even number.
            this.endingAddressRegister = endingAddressCoil;
        }

        public boolean matchesCoil(OmronFinsTag OmronFinsTag) {
            int tagStartingAddress = OmronFinsTag.getAddrword();
            int tagEndingAddress = OmronFinsTag.getAddrword() + OmronFinsTag.getNumberOfElements();
            return ((tagStartingAddress >= this.startingAddress) && (tagEndingAddress <= this.endingAddressCoil));
        }

        public boolean matchesRegister(OmronFinsTag OmronFinsTag) {
            int tagStartingAddress = OmronFinsTag.getAddrword();
            int tagEndingAddress = OmronFinsTag.getAddrword() + (int) Math.ceil((double) OmronFinsTag.getLengthBytes() / 2.0f);
            return ((tagStartingAddress >= this.startingAddress) && (tagEndingAddress <= this.endingAddressRegister));
        }

        public PlcResponseCode getResponseCode() {
            return responseCode;
        }

        public byte[] getResponseData() {
            return responseData;
        }

        public byte[] getResponseDataForTag(OmronFinsTag OmronFinsTag) {
            byte[] itemData = new byte[OmronFinsTag.getLengthBytes()];
            System.arraycopy(responseData, (OmronFinsTag.getAddrword() - startingAddress) * 2, itemData, 0, OmronFinsTag.getLengthBytes());
            return itemData;
        }
    }

    protected interface TagFactory {
        PlcTag createTag(int address, int count, OmronFinsDataType dataType);
    }



}
