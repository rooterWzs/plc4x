package org.apache.plc4x;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Collection;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Main {


    private static final Logger LOGGER = LoggerFactory.getLogger(Main.class);


    public static void main(String[] args) {
        try{
            PlcConnection connection = new DefaultPlcDriverManager().getConnection("omron-fins:tcp://127.0.0.1:9600");
            if (!connection.isConnected()){
                System.out.println("disconnected");
            }
            System.out.println("connected");

            final PlcReadRequest.Builder readrequest = connection.readRequestBuilder(); //(2.2)
            readrequest.addTagAddress("DM100", "DM100:INT[2]"); //(3.1)
//            readrequest.addTagAddress("DM400", "DM400:INT"); //(3.1)
//            readrequest.addTagAddress("DM800", "DM800:INT"); //(3.1)
//            readrequest.addTagAddress("HR110", "DM110:INT"); //(3.1)

            final PlcReadRequest rr = readrequest.build(); //(3.2)
            final PlcReadResponse response = rr.execute().get(); //(3.3)
            if (response.getResponseCode("DM100") == PlcResponseCode.OK) {//(3.4)
                System.out.println("response: " + response.getObject("DM100"));
            }
        }catch (Exception ex){
            ex.printStackTrace();
        }
    }

    public static void main2(String[] args) {

        try (PlcConnection connection = new DefaultPlcDriverManager().getConnection("modbus-rtu:tcp://127.0.0.1:502")) {
            PlcReadRequest request = connection.readRequestBuilder()
                .addTagAddress("tempSensor", "40001:UINT[1]")
                .build();
            PlcReadResponse response = request.execute().get();
            System.out.println("温度值：" + response.getFloat("tempSensor"));
        } catch (Exception e) {
            System.err.println("读取失败：" + e.getMessage());
        }

    }

}