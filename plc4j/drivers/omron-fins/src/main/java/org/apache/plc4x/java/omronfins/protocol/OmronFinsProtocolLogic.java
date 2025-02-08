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
import org.apache.plc4x.java.omronfins.configuration.OmronFinsConfiguration;
import org.apache.plc4x.java.omronfins.readwrite.FinsHandshakeRequest;
import org.apache.plc4x.java.omronfins.readwrite.FinsHandshakeResponse;
import org.apache.plc4x.java.omronfins.readwrite.FinsMessage;
import org.apache.plc4x.java.omronfins.tag.OmronFinsTag;
import org.apache.plc4x.java.omronfins.tag.OmronFinsTagHandler;
import org.apache.plc4x.java.omronfins.utils.MakeDataUtil;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.connection.PlcTagHandler;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.BiConsumer;
import java.util.function.Consumer;

public class OmronFinsProtocolLogic extends Plc4xProtocolBase<FinsMessage> implements HasConfiguration<OmronFinsConfiguration> {

    private static final Logger logger = LoggerFactory.getLogger(OmronFinsProtocolLogic.class);
    public static final Duration REQUEST_TIMEOUT = Duration.ofMillis(10000);

    private OmronFinsConfiguration configuration;

    private RequestTransactionManager tm;

    @Override
    public void setConfiguration(OmronFinsConfiguration configuration) {
        System.out.println("Connection setConfiguration:");
        this.configuration = configuration;
        // Set the transaction manager to allow only one message at a time.
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

    @Override
    public void onDisconnect(ConversationContext<FinsMessage> context) {
        // Intentionally do nothing here
        System.out.println("Connection onDisconnect:");
    }

    @Override
    public void onDiscover(ConversationContext<FinsMessage> context) {
        // Intentionally do nothing here
        System.out.println("Connection onDiscover:");
    }

//    46494e53  00 00 00 10
//    00000000  00 00 00 00
//    00000004  00 00 00 00
    @Override
    public void onConnect(ConversationContext<FinsMessage> context) {

        logger.debug("Sending Connection Request");
        FinsMessage connectionRequest = new FinsMessage(0x00000000, createFinsHandshakeRequest());

        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        context.sendRequest(connectionRequest)
            .onTimeout(e -> {
                logger.warn("Timeout during Connection establishing, closing channel...");
                context.getChannel().close();
            })
            .expectResponse(FinsMessage.class, REQUEST_TIMEOUT)
            .unwrap(FinsMessage::getMessageBody)
            .only(FinsHandshakeResponse.class)
            .check(p -> true)
            .handle(handshakeResponse -> {
                System.out.println("FinsHandshakeResponse.getServerNode():");
                // sessionHandle = FinsHandshakeResponse.getSessionHandle();
                // Send an event that connection setup is complete.
                context.fireConnected();
            });
    }

    private FinsHandshakeRequest createFinsHandshakeRequest() {
        return new FinsHandshakeRequest(MakeDataUtil.getSA1());
    }


    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        System.out.println("Connection read:");
        // TODO: Warning ... we are sending one request per tag ...
        //  the result has to be merged back together ...
        for (String tagName : readRequest.getTagNames()) {
            PlcTag tag = readRequest.getTag(tagName);
            if (!(tag instanceof OmronFinsTag)) {
                logger.error("The tag should have been of type OmronFinsTag");
            }
            OmronFinsTag omronFinsTag = (OmronFinsTag) tag;

        }
        // TODO: Should return an aggregated future ....
        return null;
    }

    @Override
    protected void decode(ConversationContext<FinsMessage> context, FinsMessage msg) throws Exception {
        System.out.println("Connection decode:");

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
