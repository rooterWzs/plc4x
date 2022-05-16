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

// IdentifyReplyCommandLogicalAssignment is the data-structure of this message
type IdentifyReplyCommandLogicalAssignment struct {
	*IdentifyReplyCommand
}

// IIdentifyReplyCommandLogicalAssignment is the corresponding interface of IdentifyReplyCommandLogicalAssignment
type IIdentifyReplyCommandLogicalAssignment interface {
	IIdentifyReplyCommand
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

func (m *IdentifyReplyCommandLogicalAssignment) GetAttribute() Attribute {
	return Attribute_LogicalAssignment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *IdentifyReplyCommandLogicalAssignment) InitializeParent(parent *IdentifyReplyCommand) {}

func (m *IdentifyReplyCommandLogicalAssignment) GetParent() *IdentifyReplyCommand {
	return m.IdentifyReplyCommand
}

// NewIdentifyReplyCommandLogicalAssignment factory function for IdentifyReplyCommandLogicalAssignment
func NewIdentifyReplyCommandLogicalAssignment() *IdentifyReplyCommandLogicalAssignment {
	_result := &IdentifyReplyCommandLogicalAssignment{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	_result.Child = _result
	return _result
}

func CastIdentifyReplyCommandLogicalAssignment(structType interface{}) *IdentifyReplyCommandLogicalAssignment {
	if casted, ok := structType.(IdentifyReplyCommandLogicalAssignment); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandLogicalAssignment); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandLogicalAssignment(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandLogicalAssignment(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandLogicalAssignment) GetTypeName() string {
	return "IdentifyReplyCommandLogicalAssignment"
}

func (m *IdentifyReplyCommandLogicalAssignment) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandLogicalAssignment) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandLogicalAssignment) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandLogicalAssignmentParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommandLogicalAssignment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandLogicalAssignment"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandLogicalAssignment"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandLogicalAssignment{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child, nil
}

func (m *IdentifyReplyCommandLogicalAssignment) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandLogicalAssignment"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandLogicalAssignment"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandLogicalAssignment) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
