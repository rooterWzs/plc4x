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
package org.apache.plc4x.java.omronfins.configuration;

import org.apache.plc4x.java.spi.configuration.PlcConnectionConfiguration;
import org.apache.plc4x.java.spi.configuration.annotations.ConfigurationParameter;
import org.apache.plc4x.java.spi.configuration.annotations.Description;
import org.apache.plc4x.java.spi.configuration.annotations.defaults.BooleanDefaultValue;
import org.apache.plc4x.java.spi.generation.ByteOrder;

public class OmronFinsConfiguration implements PlcConnectionConfiguration {


    @ConfigurationParameter("big-endian")
    @BooleanDefaultValue(true)
    @Description("Configure if the connection should be set to transport data in Big-Endian format, or not.")
    private boolean bigEndian = true;


    public ByteOrder getByteOrder() {
        return this.bigEndian ? ByteOrder.BIG_ENDIAN : ByteOrder.LITTLE_ENDIAN;
    }

    public void setByteOrder(ByteOrder byteOrder) {
        this.bigEndian = byteOrder == ByteOrder.BIG_ENDIAN;
    }



    @ConfigurationParameter("sa1")
    @Description("Configure the SA1")
    public short sa1;

    public short getSa1() {
        return sa1;
    }

    public void setSa1(short sa1) {
        this.sa1 = sa1;
    }

    @ConfigurationParameter("da1")
    @Description("Configure the DA1")
    public short da1;

    public short getDa1() {
        return da1;
    }

    public void setDa1(short da1) {
        this.da1 = da1;
    }


}
