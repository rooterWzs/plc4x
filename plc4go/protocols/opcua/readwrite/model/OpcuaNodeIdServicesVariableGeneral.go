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

// OpcuaNodeIdServicesVariableGeneral is an enum
type OpcuaNodeIdServicesVariableGeneral int32

type IOpcuaNodeIdServicesVariableGeneral interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Changes               OpcuaNodeIdServicesVariableGeneral = 2134
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventId               OpcuaNodeIdServicesVariableGeneral = 3680
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventType             OpcuaNodeIdServicesVariableGeneral = 3681
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceNode            OpcuaNodeIdServicesVariableGeneral = 3682
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceName            OpcuaNodeIdServicesVariableGeneral = 3683
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Time                  OpcuaNodeIdServicesVariableGeneral = 3684
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ReceiveTime           OpcuaNodeIdServicesVariableGeneral = 3685
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_LocalTime             OpcuaNodeIdServicesVariableGeneral = 3686
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Message               OpcuaNodeIdServicesVariableGeneral = 3687
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Severity              OpcuaNodeIdServicesVariableGeneral = 3688
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassId      OpcuaNodeIdServicesVariableGeneral = 31891
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassName    OpcuaNodeIdServicesVariableGeneral = 31892
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassId   OpcuaNodeIdServicesVariableGeneral = 31893
	OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassName OpcuaNodeIdServicesVariableGeneral = 31894
)

var OpcuaNodeIdServicesVariableGeneralValues []OpcuaNodeIdServicesVariableGeneral

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableGeneralValues = []OpcuaNodeIdServicesVariableGeneral{
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Changes,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventId,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventType,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceNode,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceName,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Time,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ReceiveTime,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_LocalTime,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Message,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Severity,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassId,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassName,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassId,
		OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassName,
	}
}

func OpcuaNodeIdServicesVariableGeneralByValue(value int32) (enum OpcuaNodeIdServicesVariableGeneral, ok bool) {
	switch value {
	case 2134:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Changes, true
	case 31891:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassId, true
	case 31892:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassName, true
	case 31893:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassId, true
	case 31894:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassName, true
	case 3680:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventId, true
	case 3681:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventType, true
	case 3682:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceNode, true
	case 3683:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceName, true
	case 3684:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Time, true
	case 3685:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ReceiveTime, true
	case 3686:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_LocalTime, true
	case 3687:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Message, true
	case 3688:
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableGeneralByName(value string) (enum OpcuaNodeIdServicesVariableGeneral, ok bool) {
	switch value {
	case "GeneralModelChangeEventType_Changes":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Changes, true
	case "GeneralModelChangeEventType_ConditionClassId":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassId, true
	case "GeneralModelChangeEventType_ConditionClassName":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassName, true
	case "GeneralModelChangeEventType_ConditionSubClassId":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassId, true
	case "GeneralModelChangeEventType_ConditionSubClassName":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassName, true
	case "GeneralModelChangeEventType_EventId":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventId, true
	case "GeneralModelChangeEventType_EventType":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventType, true
	case "GeneralModelChangeEventType_SourceNode":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceNode, true
	case "GeneralModelChangeEventType_SourceName":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceName, true
	case "GeneralModelChangeEventType_Time":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Time, true
	case "GeneralModelChangeEventType_ReceiveTime":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ReceiveTime, true
	case "GeneralModelChangeEventType_LocalTime":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_LocalTime, true
	case "GeneralModelChangeEventType_Message":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Message, true
	case "GeneralModelChangeEventType_Severity":
		return OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Severity, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableGeneralKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableGeneralValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableGeneral(structType any) OpcuaNodeIdServicesVariableGeneral {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableGeneral {
		if sOpcuaNodeIdServicesVariableGeneral, ok := typ.(OpcuaNodeIdServicesVariableGeneral); ok {
			return sOpcuaNodeIdServicesVariableGeneral
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableGeneral) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableGeneral) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableGeneralParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableGeneral, error) {
	return OpcuaNodeIdServicesVariableGeneralParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableGeneralParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableGeneral, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableGeneral", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableGeneral")
	}
	if enum, ok := OpcuaNodeIdServicesVariableGeneralByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableGeneral")
		return OpcuaNodeIdServicesVariableGeneral(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableGeneral) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableGeneral) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableGeneral", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableGeneral) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Changes:
		return "GeneralModelChangeEventType_Changes"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassId:
		return "GeneralModelChangeEventType_ConditionClassId"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionClassName:
		return "GeneralModelChangeEventType_ConditionClassName"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassId:
		return "GeneralModelChangeEventType_ConditionSubClassId"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ConditionSubClassName:
		return "GeneralModelChangeEventType_ConditionSubClassName"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventId:
		return "GeneralModelChangeEventType_EventId"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_EventType:
		return "GeneralModelChangeEventType_EventType"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceNode:
		return "GeneralModelChangeEventType_SourceNode"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_SourceName:
		return "GeneralModelChangeEventType_SourceName"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Time:
		return "GeneralModelChangeEventType_Time"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_ReceiveTime:
		return "GeneralModelChangeEventType_ReceiveTime"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_LocalTime:
		return "GeneralModelChangeEventType_LocalTime"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Message:
		return "GeneralModelChangeEventType_Message"
	case OpcuaNodeIdServicesVariableGeneral_GeneralModelChangeEventType_Severity:
		return "GeneralModelChangeEventType_Severity"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableGeneral) String() string {
	return e.PLC4XEnumName()
}
