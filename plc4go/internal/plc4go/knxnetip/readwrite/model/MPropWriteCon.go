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
)

// Code generated by code-generation. DO NOT EDIT.

// MPropWriteCon is the data-structure of this message
type MPropWriteCon struct {
	*CEMI

	// Arguments.
	Size uint16
}

// IMPropWriteCon is the corresponding interface of MPropWriteCon
type IMPropWriteCon interface {
	ICEMI
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *MPropWriteCon) GetMessageCode() uint8 {
	return 0xF5
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *MPropWriteCon) InitializeParent(parent *CEMI) {}

func (m *MPropWriteCon) GetParent() *CEMI {
	return m.CEMI
}

// NewMPropWriteCon factory function for MPropWriteCon
func NewMPropWriteCon(size uint16) *MPropWriteCon {
	_result := &MPropWriteCon{
		CEMI: NewCEMI(size),
	}
	_result.Child = _result
	return _result
}

func CastMPropWriteCon(structType interface{}) *MPropWriteCon {
	if casted, ok := structType.(MPropWriteCon); ok {
		return &casted
	}
	if casted, ok := structType.(*MPropWriteCon); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastMPropWriteCon(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastMPropWriteCon(casted.Child)
	}
	return nil
}

func (m *MPropWriteCon) GetTypeName() string {
	return "MPropWriteCon"
}

func (m *MPropWriteCon) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MPropWriteCon) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *MPropWriteCon) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MPropWriteConParse(readBuffer utils.ReadBuffer, size uint16) (*MPropWriteCon, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropWriteCon"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteCon"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MPropWriteCon{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child, nil
}

func (m *MPropWriteCon) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteCon"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("MPropWriteCon"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MPropWriteCon) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
