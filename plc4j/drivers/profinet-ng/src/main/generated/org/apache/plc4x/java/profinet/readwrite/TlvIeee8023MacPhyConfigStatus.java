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
package org.apache.plc4x.java.profinet.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class TlvIeee8023MacPhyConfigStatus extends TlvOrgSpecificIeee8023Unit implements Message {

  // Accessors for discriminator values.
  public TlvIEEESubType getSubType() {
    return TlvIEEESubType.MAC_PHY_CONFIG_STATUS;
  }

  // Properties.
  protected final short negotiationSupport;
  protected final int negotiationCapability;
  protected final int operationalMauType;

  public TlvIeee8023MacPhyConfigStatus(
      short negotiationSupport, int negotiationCapability, int operationalMauType) {
    super();
    this.negotiationSupport = negotiationSupport;
    this.negotiationCapability = negotiationCapability;
    this.operationalMauType = operationalMauType;
  }

  public short getNegotiationSupport() {
    return negotiationSupport;
  }

  public int getNegotiationCapability() {
    return negotiationCapability;
  }

  public int getOperationalMauType() {
    return operationalMauType;
  }

  @Override
  protected void serializeTlvOrgSpecificIeee8023UnitChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("TlvIeee8023MacPhyConfigStatus");

    // Simple Field (negotiationSupport)
    writeSimpleField("negotiationSupport", negotiationSupport, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (negotiationCapability)
    writeSimpleField(
        "negotiationCapability", negotiationCapability, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (operationalMauType)
    writeSimpleField("operationalMauType", operationalMauType, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("TlvIeee8023MacPhyConfigStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    TlvIeee8023MacPhyConfigStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (negotiationSupport)
    lengthInBits += 8;

    // Simple field (negotiationCapability)
    lengthInBits += 16;

    // Simple field (operationalMauType)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static TlvOrgSpecificIeee8023UnitBuilder staticParseTlvOrgSpecificIeee8023UnitBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("TlvIeee8023MacPhyConfigStatus");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short negotiationSupport =
        readSimpleField("negotiationSupport", readUnsignedShort(readBuffer, 8));

    int negotiationCapability =
        readSimpleField("negotiationCapability", readUnsignedInt(readBuffer, 16));

    int operationalMauType = readSimpleField("operationalMauType", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("TlvIeee8023MacPhyConfigStatus");
    // Create the instance
    return new TlvIeee8023MacPhyConfigStatusBuilderImpl(
        negotiationSupport, negotiationCapability, operationalMauType);
  }

  public static class TlvIeee8023MacPhyConfigStatusBuilderImpl
      implements TlvOrgSpecificIeee8023Unit.TlvOrgSpecificIeee8023UnitBuilder {
    private final short negotiationSupport;
    private final int negotiationCapability;
    private final int operationalMauType;

    public TlvIeee8023MacPhyConfigStatusBuilderImpl(
        short negotiationSupport, int negotiationCapability, int operationalMauType) {
      this.negotiationSupport = negotiationSupport;
      this.negotiationCapability = negotiationCapability;
      this.operationalMauType = operationalMauType;
    }

    public TlvIeee8023MacPhyConfigStatus build() {
      TlvIeee8023MacPhyConfigStatus tlvIeee8023MacPhyConfigStatus =
          new TlvIeee8023MacPhyConfigStatus(
              negotiationSupport, negotiationCapability, operationalMauType);
      return tlvIeee8023MacPhyConfigStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof TlvIeee8023MacPhyConfigStatus)) {
      return false;
    }
    TlvIeee8023MacPhyConfigStatus that = (TlvIeee8023MacPhyConfigStatus) o;
    return (getNegotiationSupport() == that.getNegotiationSupport())
        && (getNegotiationCapability() == that.getNegotiationCapability())
        && (getOperationalMauType() == that.getOperationalMauType())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNegotiationSupport(),
        getNegotiationCapability(),
        getOperationalMauType());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}