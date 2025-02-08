/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package org.apache.plc4x.java.omronfins.context;

import org.apache.plc4x.java.omronfins.configuration.OmronFinsConfiguration;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.context.DriverContext;

public class OmronFinsDriverContext implements DriverContext, HasConfiguration<OmronFinsConfiguration> {

    private short sa1;

    private short da1;


    @Override
    public void setConfiguration(OmronFinsConfiguration configuration) {
        this.sa1 = configuration.getSa1();
        this.da1 = configuration.getDa1();
    }


    public short getDa1() {
        return da1;
    }

    public void setDa1(short da1) {
        this.da1 = da1;
    }

    public short getSa1() {
        return sa1;
    }

    public void setSa1(short sa1) {
        this.sa1 = sa1;
    }

}
