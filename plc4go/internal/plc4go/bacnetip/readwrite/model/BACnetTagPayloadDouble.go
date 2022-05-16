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

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetTagPayloadDouble is the data-structure of this message
type BACnetTagPayloadDouble struct {
	Value float64
}

// IBACnetTagPayloadDouble is the corresponding interface of BACnetTagPayloadDouble
type IBACnetTagPayloadDouble interface {
	// GetValue returns Value (property field)
	GetValue() float64
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetTagPayloadDouble) GetValue() float64 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadDouble factory function for BACnetTagPayloadDouble
func NewBACnetTagPayloadDouble(value float64) *BACnetTagPayloadDouble {
	return &BACnetTagPayloadDouble{Value: value}
}

func CastBACnetTagPayloadDouble(structType interface{}) *BACnetTagPayloadDouble {
	if casted, ok := structType.(BACnetTagPayloadDouble); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagPayloadDouble); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagPayloadDouble) GetTypeName() string {
	return "BACnetTagPayloadDouble"
}

func (m *BACnetTagPayloadDouble) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadDouble) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 64

	return lengthInBits
}

func (m *BACnetTagPayloadDouble) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadDoubleParse(readBuffer utils.ReadBuffer) (*BACnetTagPayloadDouble, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadDouble"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadFloat64("value", 64)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadDouble"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadDouble(value), nil
}

func (m *BACnetTagPayloadDouble) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadDouble"); pushErr != nil {
		return pushErr
	}

	// Simple Field (value)
	value := float64(m.Value)
	_valueErr := writeBuffer.WriteFloat64("value", 64, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadDouble"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
