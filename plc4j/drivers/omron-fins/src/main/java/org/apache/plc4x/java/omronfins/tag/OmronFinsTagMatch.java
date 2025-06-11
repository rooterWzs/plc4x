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
package org.apache.plc4x.java.omronfins.tag;

import org.apache.plc4x.java.api.exceptions.PlcInvalidTagException;
import org.apache.plc4x.java.omronfins.readwrite.FinsRegisterAddress;
import org.apache.plc4x.java.omronfins.readwrite.OmronFinsDataType;
import org.apache.plc4x.java.spi.tag.TagConfigParser;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class OmronFinsTagMatch extends OmronFinsTag {

    public static final String ADDRESS_PREFIX = "0x";

//    DM100:INT
    private static final List<Pattern> ADDRESS_PATTERNS = Arrays.asList(
        Pattern.compile(
        "^" +
            "(?<addrprefix>DM|HR|WR|CIO|EM)" +     // 地址前缀
            "(?<addrword>\\d+)" +                  // 主地址数字
            "(?:\\.(?<addrbit>\\d+))?" +           // 可选：子地址位
            ":" +
            "(?<datatype>USINT|SINT|INT|UINT|STRING|BOOL)" + // 数据类型
            "(?:\\[(?<quantity>\\d+)\\])?" +       // 可选：数组数量
            "$"
        )
    );

    protected static final int REGISTER_MAX_ADDRESS = 65535;

    public OmronFinsTagMatch(FinsRegisterAddress addrprefix, int addrword, byte addrbit, Integer quantity, OmronFinsDataType dataType, Map<String, String> config) {
        super(addrprefix, addrword, addrbit, quantity, dataType, config);
    }

    protected String getAddressStringPrefix() {
        return ADDRESS_PREFIX;
    }

    @Override
    public int getLogicalAddress() {
        return getAddrword() + PROTOCOL_ADDRESS_OFFSET;
    }

    public static boolean matches(String addressString) {
        return ADDRESS_PATTERNS.stream().anyMatch(p -> p.matcher(addressString).matches());
    }

    private static Optional<Matcher> findMatchingPattern(String addressString) {
        return ADDRESS_PATTERNS.stream()
            .map(p -> p.matcher(addressString))
            .filter(Matcher::matches)
            .findFirst();
    }

    public static OmronFinsTagMatch of(String addressString) {
        Matcher matcher = findMatchingPattern(addressString)
            .orElseThrow(() -> new PlcInvalidTagException(addressString, ADDRESS_PATTERNS.get(0)));

        if (matcher.matches()) {
            System.out.println("addr_prefix: " + matcher.group("addrprefix"));  // DM
            System.out.println("address_word: " + matcher.group("addrword")); // 100
            System.out.println("address_bit: " + (matcher.group("addrbit") != null ? matcher.group("addrbit") : "")); // 无
            System.out.println("data_type: " + matcher.group("datatype"));       // INT
            System.out.println("quantity: " + matcher.group("quantity"));       // INT
        }

        FinsRegisterAddress addrprefix = Optional.ofNullable(matcher.group("addrprefix"))
            .map(FinsRegisterAddress::valueOf)
            .orElse(FinsRegisterAddress.DM);

        int addrword = Integer.parseInt(matcher.group("addrword")) - PROTOCOL_ADDRESS_OFFSET;
        if (addrword > REGISTER_MAX_ADDRESS) {
            throw new IllegalArgumentException("Address must be ≤ " + REGISTER_MAX_ADDRESS + ". Got " + (addrword + PROTOCOL_ADDRESS_OFFSET));
        }

        byte addrbit = Optional.ofNullable(matcher.group("addrbit")).map(Byte::parseByte).orElse((byte) 0);
        if (addrbit < 0 || addrbit > 15) {
            throw new IllegalArgumentException("Address must >=0 and be ≤ 15.");
        }


        int quantity = Optional.ofNullable(matcher.group("quantity")).map(Integer::parseInt).orElse(1);
        if (addrword + quantity > REGISTER_MAX_ADDRESS) {
            throw new IllegalArgumentException("Out of range: final address is " + (addrword + PROTOCOL_ADDRESS_OFFSET + quantity - 1));
        }
        if (quantity > 2000) {
            throw new IllegalArgumentException("Quantity cannot exceed 2000. Was " + quantity);
        }

        OmronFinsDataType dataType = Optional.ofNullable(matcher.group("datatype"))
            .map(OmronFinsDataType::valueOf)
            .orElse(OmronFinsDataType.BOOL);

        Map<String, String> config = TagConfigParser.parse(addressString);

        return new OmronFinsTagMatch(addrprefix, addrword, addrbit, quantity, dataType, config);
    }
}

