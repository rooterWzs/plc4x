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

public class VariantInt16 extends Variant implements Message {

  // Accessors for discriminator values.
  public Short getVariantType() {
    return (short) 4;
  }

  // Properties.
  protected final Integer arrayLength;
  protected final List<Short> value;

  public VariantInt16(
      boolean arrayLengthSpecified,
      boolean arrayDimensionsSpecified,
      Integer noOfArrayDimensions,
      List<Boolean> arrayDimensions,
      Integer arrayLength,
      List<Short> value) {
    super(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions);
    this.arrayLength = arrayLength;
    this.value = value;
  }

  public Integer getArrayLength() {
    return arrayLength;
  }

  public List<Short> getValue() {
    return value;
  }

  @Override
  protected void serializeVariantChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("VariantInt16");

    // Optional Field (arrayLength) (Can be skipped, if the value is null)
    writeOptionalField(
        "arrayLength", arrayLength, writeSignedInt(writeBuffer, 32), arrayLengthSpecified);

    // Array Field (value)
    writeSimpleTypeArrayField("value", value, writeSignedShort(writeBuffer, 16));

    writeBuffer.popContext("VariantInt16");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    VariantInt16 _value = this;

    // Optional Field (arrayLength)
    if (arrayLength != null) {
      lengthInBits += 32;
    }

    // Array field
    if (value != null) {
      lengthInBits += 16 * value.size();
    }

    return lengthInBits;
  }

  public static VariantInt16Builder staticParseBuilder(
      ReadBuffer readBuffer, Boolean arrayLengthSpecified) throws ParseException {
    readBuffer.pullContext("VariantInt16");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    Integer arrayLength =
        readOptionalField("arrayLength", readSignedInt(readBuffer, 32), arrayLengthSpecified);

    List<Short> value =
        readCountArrayField(
            "value",
            readSignedShort(readBuffer, 16),
            (((arrayLength) == (null)) ? 1 : arrayLength));

    readBuffer.closeContext("VariantInt16");
    // Create the instance
    return new VariantInt16Builder(arrayLength, value);
  }

  public static class VariantInt16Builder implements Variant.VariantBuilder {
    private final Integer arrayLength;
    private final List<Short> value;

    public VariantInt16Builder(Integer arrayLength, List<Short> value) {

      this.arrayLength = arrayLength;
      this.value = value;
    }

    public VariantInt16 build(
        boolean arrayLengthSpecified,
        boolean arrayDimensionsSpecified,
        Integer noOfArrayDimensions,
        List<Boolean> arrayDimensions) {
      VariantInt16 variantInt16 =
          new VariantInt16(
              arrayLengthSpecified,
              arrayDimensionsSpecified,
              noOfArrayDimensions,
              arrayDimensions,
              arrayLength,
              value);
      return variantInt16;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof VariantInt16)) {
      return false;
    }
    VariantInt16 that = (VariantInt16) o;
    return (getArrayLength() == that.getArrayLength())
        && (getValue() == that.getValue())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getArrayLength(), getValue());
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