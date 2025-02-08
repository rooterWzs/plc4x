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
package org.apache.plc4x.java.omronfins;

import io.netty.buffer.ByteBuf;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.omronfins.configuration.OmronFinsConfiguration;
import org.apache.plc4x.java.omronfins.configuration.OmronFinsTcpTransportConfiguration;
import org.apache.plc4x.java.omronfins.context.OmronFinsDriverContext;
import org.apache.plc4x.java.omronfins.protocol.OmronFinsProtocolLogic;
import org.apache.plc4x.java.omronfins.readwrite.FinsMessage;
import org.apache.plc4x.java.omronfins.tag.OmronFinsTag;
import org.apache.plc4x.java.spi.configuration.PlcConnectionConfiguration;
import org.apache.plc4x.java.spi.configuration.PlcTransportConfiguration;
import org.apache.plc4x.java.spi.connection.GeneratedDriverBase;
import org.apache.plc4x.java.spi.connection.ProtocolStackConfigurer;
import org.apache.plc4x.java.spi.connection.SingleProtocolStackConfigurer;

import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.function.ToIntFunction;

public class OmronFinsDriver extends GeneratedDriverBase<FinsMessage> {

    public static final int OMRON_FINS_PORT = 9600;

    private OmronFinsConfiguration configuration;

    @Override
    public String getProtocolCode() {
        return "omron-fins";
    }

    @Override
    public String getProtocolName() {
        return "OMRON FINS";
    }


    @Override
    protected Class<? extends PlcConnectionConfiguration> getConfigurationClass() {
        return OmronFinsConfiguration.class;
    }

    @Override
    protected Optional<Class<? extends PlcTransportConfiguration>> getTransportConfigurationClass(String transportCode) {
        switch (transportCode) {
            case "tcp":
                return Optional.of(OmronFinsTcpTransportConfiguration.class);
        }
        return Optional.empty();
    }

    @Override
    protected Optional<String> getDefaultTransportCode() {
        return Optional.of("raw");
    }

    @Override
    protected List<String> getSupportedTransportCodes() {
        return Collections.singletonList("tcp");
    }

    /**
     * This protocol doesn't have a disconnect procedure, so there is no need to wait for a login to finish.
     * @return false
     */
    @Override
    protected boolean awaitDisconnectComplete() {
        return false;
    }

    @Override
    protected boolean canRead() {
        return true;
    }

    @Override
    protected boolean canWrite() {
        return true;
    }


    @Override
    protected ProtocolStackConfigurer<FinsMessage> getStackConfigurer() {
        return SingleProtocolStackConfigurer.builder(FinsMessage.class, (io) -> FinsMessage.staticParse(io, false))
            .withProtocol(OmronFinsProtocolLogic.class)
            .withDriverContext(OmronFinsDriverContext.class)
//            .withPacketSizeEstimator(ByteLengthEstimator.class)
//            .byteOrder(configuration.getByteOrder())
            .build();
    }


    /**
     * Estimate the Length of a Packet
     */
    public static class ByteLengthEstimator implements ToIntFunction<ByteBuf> {
        @Override
        public int applyAsInt(ByteBuf byteBuf) {
            if (byteBuf.readableBytes() >= 4) {
                // In the mspec we subtract 28 from the full size ... so here we gotta add it back.
                return byteBuf.getUnsignedShort(byteBuf.readerIndex() + 2) + 28;
            }
            return -1;
        }
    }

    @Override
    public PlcTag prepareTag(String tagAddress) {
        return OmronFinsTag.of(tagAddress);
    }

}
