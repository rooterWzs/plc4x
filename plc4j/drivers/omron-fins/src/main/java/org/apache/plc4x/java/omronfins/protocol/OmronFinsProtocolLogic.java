/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.omronfins.protocol;

import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.omronfins.configuration.OmronFinsConfiguration;
import org.apache.plc4x.java.omronfins.context.OmronFinsDriverContext;
import org.apache.plc4x.java.omronfins.readwrite.*;
import org.apache.plc4x.java.omronfins.tag.OmronFinsTag;
import org.apache.plc4x.java.omronfins.tag.OmronFinsTagHandler;
import org.apache.plc4x.java.omronfins.utils.MakeDataUtil;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.connection.PlcTagHandler;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadRequest;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadResponse;
import org.apache.plc4x.java.spi.messages.utils.DefaultPlcResponseItem;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.crypto.spec.PSource;
import java.time.Duration;
import java.util.Collections;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.BiConsumer;
import java.util.function.Consumer;

public class OmronFinsProtocolLogic extends Plc4xProtocolBase<FinsMessage> {

    private static final Logger logger = LoggerFactory.getLogger(OmronFinsProtocolLogic.class);
    public static final Duration REQUEST_TIMEOUT = Duration.ofMillis(10000);

    private OmronFinsConfiguration configuration;

    private RequestTransactionManager tm;

    private OmronFinsDriverContext driverContext;

    private AtomicInteger transactionId = new AtomicInteger(0);

    @Override
    public void setDriverContext(DriverContext driverContext) {
        super.setDriverContext(driverContext);
        this.driverContext = (OmronFinsDriverContext) driverContext;
        this.tm = new RequestTransactionManager(1);
    }

    @Override
    public PlcTagHandler getTagHandler() {
        System.out.println("Connection getTagHandler:");
        return new OmronFinsTagHandler();
    }

    @Override
    public void close(ConversationContext<FinsMessage> context) {
        System.out.println("Connection close:");
        tm.shutdown();
    }

//    @Override
//    public void onDisconnect(ConversationContext<FinsMessage> context) {
//        // Intentionally do nothing here
//        System.out.println("Connection onDisconnect:");
//    }

//    @Override
//    public void onDiscover(ConversationContext<FinsMessage> context) {
//        // Intentionally do nothing here
//        System.out.println("Connection onDiscover:");
//    }

//    46494e53  00 00 00 10
//    00000000  00 00 00 00
//    00000004  00 00 00 00
    @Override
    public void onConnect(ConversationContext<FinsMessage> context) {
        FinsMessage connectionRequest = new FinsMessage(0x00000000, createFinsHandshakeRequest());
        logger.debug("Sending Connection Request：" + connectionRequest.toString());
        context.sendRequest(connectionRequest)
            .onTimeout(e -> {
                logger.warn("Timeout during Connection establishing, closing channel...");
//                context.getChannel().close();
            })
            .expectResponse(FinsMessage.class, REQUEST_TIMEOUT)
            .unwrap(FinsMessage::getMessageBody)
            .only(FinsHandshakeResponse.class)
            .handle(handshakeResponse -> {
                System.out.println("FinsHandshakeResponse.getServerNode():"+ handshakeResponse.getSA1());
                // sessionHandle = FinsHandshakeResponse.getSessionHandle();
                // Send an event that connection setup is complete.
                OmronFinsConfiguration configuration = new OmronFinsConfiguration();
                configuration.setSa1(handshakeResponse.getSA1());
                configuration.setDa1(handshakeResponse.getDA1());
                driverContext.setConfiguration(configuration);
                context.fireConnected();
            });
    }

    private FinsHandshakeRequest createFinsHandshakeRequest() {
        return new FinsHandshakeRequest(MakeDataUtil.getSA1());
    }


    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();
        DefaultPlcReadRequest request = (DefaultPlcReadRequest) readRequest;

        System.out.println("Connection read: SA1:" + driverContext.getSa1() + " | DA1:" + driverContext.getDa1());

        // TODO: Warning ... we are sending one request per tag ...
        //  the result has to be merged back together ...
        for (String tagName : readRequest.getTagNames()) {
            PlcTag tag = readRequest.getTag(tagName);
            if (!(tag instanceof OmronFinsTag)) {
                logger.error("The tag should have been of type OmronFinsTag");
            }
            OmronFinsTag omronFinsTag = (OmronFinsTag) tag;
            FinsHeaderSection section = new FinsHeaderSection((short) 0x80, (short)0x00, driverContext.getDa1(), (short)0x00,
                (short)0x00, driverContext.getSa1(), (short)0x00);


            FinsReadRequest finsReadRequest = new FinsReadRequest(section,FinsSrcMrcCode.FinsReadRequest,omronFinsTag.getAddrprefix(), omronFinsTag.getAddrword(),omronFinsTag.getAddrbit(), omronFinsTag.getQuantity());
            FinsMessage finsMessage = new FinsMessage(FinsCommandCode.FinsReadRequest.getValue(), finsReadRequest);

            RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
            transaction.submit(() -> conversationContext.sendRequest(finsMessage)
                .expectResponse(FinsMessage.class, REQUEST_TIMEOUT)
                .onTimeout(future::completeExceptionally)
//                .onError((p, e) -> future.completeExceptionally(e))
//                .check(p -> (
//                    p.getErrorCode() == FinsErrorCode.NoError.getValue()
//                ))
                .unwrap(FinsMessage::getMessageBody)
//                .only(FinsReadResponse.class)
                .handle(finsReadResponse -> {
                    System.out.println("finsReadResponse:"+ finsReadResponse.toString());
//                    if(finsReadResponse.getResErrorCode() == 0){
//
//                    }
                    System.out.println("true");
                    // Try to decode the response data based on the corresponding request.
                    PlcValue plcValue = null;
                    PlcResponseCode responseCode;
                    // Check if the response was an error response.
//                    if (responsePdu instanceof ModbusPDUError) {
//                        ModbusPDUError errorResponse = (ModbusPDUError) responsePdu;
//                        responseCode = getErrorCode(errorResponse);
//                    } else {
//                        try {
//                            ModbusByteOrder byteOrder = defaultPayloadByteOrder;
//                            if(tag.getByteOrder() != null) {
//                                byteOrder = tag.getByteOrder();
//                            }
//                            plcValue = toPlcValue(requestPdu, responsePdu, tag.getDataType(), byteOrder);
//                            responseCode = PlcResponseCode.OK;
//                        } catch (ParseException e) {
//                            // Add an error response code ...
//                            responseCode = PlcResponseCode.INTERNAL_ERROR;
//                        }
//                    }

                    // Prepare the response.
//                    PlcReadResponse response = new DefaultPlcReadResponse(request,
//                        Collections.singletonMap(tagName, new DefaultPlcResponseItem<>(responseCode, plcValue)));

                    // Pass the response back to the application.
//                    future.complete(response);

                    // Finish the request-transaction.
                    transaction.endRequest();
                }));

        }
        // TODO: Should return an aggregated future ....
        return null;
    }

    @Override
    protected void decode(ConversationContext<FinsMessage> context, FinsMessage msg) throws Exception {
        System.out.println("receive byte:" + msg);
        super.decode(context, msg);
    }




    static class TransactionErrorCallback<T, E extends Throwable> implements Consumer<TimeoutException>, BiConsumer<FinsMessage, E> {

        private final CompletableFuture<T> future;
        private final RequestTransactionManager.RequestTransaction transaction;

        TransactionErrorCallback(CompletableFuture<T> future, RequestTransactionManager.RequestTransaction transaction) {
            this.future = future;
            this.transaction = transaction;
        }

        @Override
        public void accept(TimeoutException e) {
            try {
                transaction.endRequest();
            } catch (Exception ex) {
                logger.info(ex.getMessage());
            }
            future.completeExceptionally(e);
        }

        @Override
        public void accept(FinsMessage tpktPacket, E e) {
            try {
                transaction.endRequest();
            } catch (Exception ex) {
                logger.info(ex.getMessage());
            }
            future.completeExceptionally(e);
        }
    }

}
