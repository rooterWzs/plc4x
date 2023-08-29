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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableDatagram {
  DatagramConnectionTransportType_DiscoveryAddress_NetworkInterface((int) 15154L),
  DatagramConnectionTransportType_DiscoveryAddress_NetworkInterface_Selections((int) 17579L),
  DatagramConnectionTransportType_DiscoveryAddress_NetworkInterface_SelectionDescriptions(
      (int) 17580L),
  DatagramConnectionTransportType_DiscoveryAddress_NetworkInterface_RestrictToList((int) 17581L),
  DatagramWriterGroupTransportType_MessageRepeatCount((int) 21134L),
  DatagramWriterGroupTransportType_MessageRepeatDelay((int) 21135L),
  DatagramConnectionTransportType_DiscoveryAnnounceRate((int) 23839L),
  DatagramConnectionTransportType_DiscoveryMaxMessageSize((int) 23840L),
  DatagramWriterGroupTransportType_Address_NetworkInterface((int) 23843L),
  DatagramWriterGroupTransportType_Address_NetworkInterface_Selections((int) 23844L),
  DatagramWriterGroupTransportType_Address_NetworkInterface_SelectionDescriptions((int) 23845L),
  DatagramWriterGroupTransportType_Address_NetworkInterface_RestrictToList((int) 23846L),
  DatagramWriterGroupTransportType_DatagramQos((int) 23847L),
  DatagramWriterGroupTransportType_DiscoveryAnnounceRate((int) 23848L),
  DatagramWriterGroupTransportType_Topic((int) 23849L),
  DatagramDataSetReaderTransportType_Address_NetworkInterface((int) 24018L),
  DatagramDataSetReaderTransportType_Address_NetworkInterface_Selections((int) 24019L),
  DatagramDataSetReaderTransportType_Address_NetworkInterface_SelectionDescriptions((int) 24020L),
  DatagramDataSetReaderTransportType_Address_NetworkInterface_RestrictToList((int) 24021L),
  DatagramDataSetReaderTransportType_DatagramQos((int) 24022L),
  DatagramDataSetReaderTransportType_Topic((int) 24023L),
  DatagramConnectionTransportType_QosCategory((int) 25525L),
  DatagramConnectionTransportType_DatagramQos((int) 25526L),
  DatagramWriterGroupTransportType_QosCategory((int) 25527L),
  DatagramDataSetReaderTransportType_QosCategory((int) 25528L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableDatagram> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableDatagram value : OpcuaNodeIdServicesVariableDatagram.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableDatagram(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableDatagram enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
