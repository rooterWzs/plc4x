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

public enum OpcuaNodeIdServicesVariableRemove {
  RemoveCertificateMethodType_InputArguments((int) 12521L),
  RemoveConnectionMethodType_InputArguments((int) 14184L),
  RemovePublishedDataSetMethodType_InputArguments((int) 14508L),
  RemoveSecurityGroupMethodType_InputArguments((int) 15470L),
  RemoveExtensionFieldMethodType_InputArguments((int) 15500L),
  RemoveIdentityMethodType_InputArguments((int) 15639L),
  RemoveRoleMethodType_InputArguments((int) 16006L),
  RemoveApplicationMethodType_InputArguments((int) 16187L),
  RemoveEndpointMethodType_InputArguments((int) 16191L),
  RemoveDataSetFolderMethodType_InputArguments((int) 17201L),
  RemoveSubscribedDataSetMethodType_InputArguments((int) 23825L),
  RemoveUserMethodType_InputArguments((int) 24287L),
  RemoveSecurityGroupFolderMethodType_InputArguments((int) 25292L),
  RemovePushTargetMethodType_InputArguments((int) 25380L),
  RemovePushTargetFolderMethodType_InputArguments((int) 25385L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableRemove> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableRemove value : OpcuaNodeIdServicesVariableRemove.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableRemove(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableRemove enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}