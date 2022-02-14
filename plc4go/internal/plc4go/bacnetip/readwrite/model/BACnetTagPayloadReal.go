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

// The data-structure of this message
type BACnetTagPayloadReal struct {
	Value float32
}

// The corresponding interface
type IBACnetTagPayloadReal interface {
	// GetValue returns Value
	GetValue() float32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadReal) GetValue() float32 {
	return m.Value
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadReal factory function for BACnetTagPayloadReal
func NewBACnetTagPayloadReal(value float32) *BACnetTagPayloadReal {
	return &BACnetTagPayloadReal{Value: value}
}

func CastBACnetTagPayloadReal(structType interface{}) *BACnetTagPayloadReal {
	castFunc := func(typ interface{}) *BACnetTagPayloadReal {
		if casted, ok := typ.(BACnetTagPayloadReal); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetTagPayloadReal); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetTagPayloadReal) GetTypeName() string {
	return "BACnetTagPayloadReal"
}

func (m *BACnetTagPayloadReal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadReal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 32

	return lengthInBits
}

func (m *BACnetTagPayloadReal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadRealParse(readBuffer utils.ReadBuffer) (*BACnetTagPayloadReal, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadReal"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadFloat32("value", 32)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadReal"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadReal(value), nil
}

func (m *BACnetTagPayloadReal) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadReal"); pushErr != nil {
		return pushErr
	}

	// Simple Field (value)
	value := float32(m.Value)
	_valueErr := writeBuffer.WriteFloat32("value", 32, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadReal"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadReal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}