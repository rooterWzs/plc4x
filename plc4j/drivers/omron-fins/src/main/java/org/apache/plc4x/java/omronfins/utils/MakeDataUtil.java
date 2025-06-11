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
package org.apache.plc4x.java.omronfins.utils;

import java.net.InetAddress;
import java.net.UnknownHostException;

public class MakeDataUtil {

    // 定义目标信息
    private static final byte DNA = 0x00;  // 目标网络号
    private static final byte DA2 = 0x00;  // 目标单位号（通常是0，表示 CPU单元）


    // 从IP地址获取节点号
    private static byte getNodeFromIp(String ipAddress) throws UnknownHostException {
        InetAddress inetAddress = InetAddress.getByName(ipAddress);
        byte[] ipBytes = inetAddress.getAddress();
        return ipBytes[ipBytes.length - 1];  // 获取IP地址的最后一个字节
    }


    public static Short getSA1() {
        try {
            // 获取本地主机的IP地址
            InetAddress inetAddress = InetAddress.getLocalHost();
            String pcIpAddress = inetAddress.getHostAddress();
            System.out.println("Local IP Address: " + pcIpAddress);
            String[] ipArr = pcIpAddress.split("\\.");
            return Short.valueOf(ipArr[3]);
        } catch (UnknownHostException e) {
            e.printStackTrace();
        }
        return null;
    }


    public static Long getClientNode() {
        try {
            // 获取本地主机的IP地址
            InetAddress inetAddress = InetAddress.getLocalHost();
            String pcIpAddress = inetAddress.getHostAddress();
            System.out.println("Local IP Address: " + pcIpAddress);
            byte[] clientNode = new byte[4];
            clientNode[0] = DNA;  // 网络号
            clientNode[1] = DNA;  // 网络号
            clientNode[2] = getNodeFromIp(pcIpAddress);  // 使用IP地址的最后一位作为节点号
            clientNode[3] = DA2;  // 单位号
            System.out.println("Client Node: " + byteArrayToLong(clientNode));
            return byteArrayToLong(clientNode);

        } catch (UnknownHostException e) {
            e.printStackTrace();
        }
        return null;
    }

    // 工具方法：将字节数组转换为十六进制字符串
    public static String bytesToHex(byte[] bytes) {
        StringBuilder sb = new StringBuilder();
        for (byte b : bytes) {
            sb.append(String.format("%02X", b));
        }
        return sb.toString();
    }

    // 字节数组转换为 long 类型
    // 字节数组转换为 long 类型
    public static long byteArrayToLong(byte[] byteArray) {
        long value = 0;
        for (int i = 0; i < byteArray.length; i++) {
            value |= ((long) byteArray[i] & 0xFF) << (8 * (byteArray.length - 1 - i));
        }
        return value;
    }


}
