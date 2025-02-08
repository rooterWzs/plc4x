package org.apache.plc4x;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Main {


    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);

    public static void main1(String[] args) {
        try{
//            PlcConnection connection = new DefaultPlcDriverManager().getConnection("omron-fins://192.168.31.220:9600");
            PlcConnection connection = new DefaultPlcDriverManager().getConnection("ab-eth:tcp://192.168.31.254:44818");
            if (connection.isConnected()){
                System.out.println("connected");
            }
            else{
                System.out.println("disconnected");
            }


        }catch (Exception ex){
            ex.printStackTrace();
        }


    }

    public static void main(String[] args) {
        try{
            PlcConnection connection = new DefaultPlcDriverManager().getConnection("omron-fins:tcp://127.0.0.1:9600");
            if (!connection.isConnected()){
                System.out.println("disconnected");
            }
            System.out.println("connected");

            final PlcReadRequest.Builder readrequest = connection.readRequestBuilder(); //(2.2)
            readrequest.addTagAddress("MySZL", "SZL_ID=16#0091;INDEX=16#0000"); //(3.1)
            final PlcReadRequest rr = readrequest.build(); //(3.2)
            final PlcReadResponse szlresponse = rr.execute().get(); //(3.3)
            if (szlresponse.getResponseCode("MySZL") == PlcResponseCode.OK) {//(3.4)
            }
        }catch (Exception ex){
            ex.printStackTrace();
        }


    }
}