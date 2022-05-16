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

// CommandHeader is the data-structure of this message
type CommandHeader struct {
	Value byte
}

// ICommandHeader is the corresponding interface of CommandHeader
type ICommandHeader interface {
	// GetValue returns Value (property field)
	GetValue() byte
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

func (m *CommandHeader) GetValue() byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCommandHeader factory function for CommandHeader
func NewCommandHeader(value byte) *CommandHeader {
	return &CommandHeader{Value: value}
}

func CastCommandHeader(structType interface{}) *CommandHeader {
	if casted, ok := structType.(CommandHeader); ok {
		return &casted
	}
	if casted, ok := structType.(*CommandHeader); ok {
		return casted
	}
	return nil
}

func (m *CommandHeader) GetTypeName() string {
	return "CommandHeader"
}

func (m *CommandHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CommandHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (value)
	lengthInBits += 8

	return lengthInBits
}

func (m *CommandHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CommandHeaderParse(readBuffer utils.ReadBuffer) (*CommandHeader, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CommandHeader"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (value)
	_value, _valueErr := readBuffer.ReadByte("value")
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("CommandHeader"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewCommandHeader(value), nil
}

func (m *CommandHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("CommandHeader"); pushErr != nil {
		return pushErr
	}

	// Simple Field (value)
	value := byte(m.Value)
	_valueErr := writeBuffer.WriteByte("value", (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("CommandHeader"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CommandHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
