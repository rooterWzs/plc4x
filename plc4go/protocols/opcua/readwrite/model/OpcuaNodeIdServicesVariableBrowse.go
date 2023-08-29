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

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableBrowse is an enum
type OpcuaNodeIdServicesVariableBrowse int32

type IOpcuaNodeIdServicesVariableBrowse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableBrowse_BrowseDirection_EnumStrings OpcuaNodeIdServicesVariableBrowse = 7603
	OpcuaNodeIdServicesVariableBrowse_BrowseResultMask_EnumValues OpcuaNodeIdServicesVariableBrowse = 11883
)

var OpcuaNodeIdServicesVariableBrowseValues []OpcuaNodeIdServicesVariableBrowse

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableBrowseValues = []OpcuaNodeIdServicesVariableBrowse{
		OpcuaNodeIdServicesVariableBrowse_BrowseDirection_EnumStrings,
		OpcuaNodeIdServicesVariableBrowse_BrowseResultMask_EnumValues,
	}
}

func OpcuaNodeIdServicesVariableBrowseByValue(value int32) (enum OpcuaNodeIdServicesVariableBrowse, ok bool) {
	switch value {
	case 11883:
		return OpcuaNodeIdServicesVariableBrowse_BrowseResultMask_EnumValues, true
	case 7603:
		return OpcuaNodeIdServicesVariableBrowse_BrowseDirection_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBrowseByName(value string) (enum OpcuaNodeIdServicesVariableBrowse, ok bool) {
	switch value {
	case "BrowseResultMask_EnumValues":
		return OpcuaNodeIdServicesVariableBrowse_BrowseResultMask_EnumValues, true
	case "BrowseDirection_EnumStrings":
		return OpcuaNodeIdServicesVariableBrowse_BrowseDirection_EnumStrings, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableBrowseKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableBrowseValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableBrowse(structType any) OpcuaNodeIdServicesVariableBrowse {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableBrowse {
		if sOpcuaNodeIdServicesVariableBrowse, ok := typ.(OpcuaNodeIdServicesVariableBrowse); ok {
			return sOpcuaNodeIdServicesVariableBrowse
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableBrowse) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableBrowse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableBrowseParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableBrowse, error) {
	return OpcuaNodeIdServicesVariableBrowseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableBrowseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableBrowse, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableBrowse", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableBrowse")
	}
	if enum, ok := OpcuaNodeIdServicesVariableBrowseByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableBrowse")
		return OpcuaNodeIdServicesVariableBrowse(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableBrowse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableBrowse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableBrowse", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableBrowse) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableBrowse_BrowseResultMask_EnumValues:
		return "BrowseResultMask_EnumValues"
	case OpcuaNodeIdServicesVariableBrowse_BrowseDirection_EnumStrings:
		return "BrowseDirection_EnumStrings"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableBrowse) String() string {
	return e.PLC4XEnumName()
}
