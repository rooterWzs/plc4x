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

// CALDataRequestIdentify is the data-structure of this message
type CALDataRequestIdentify struct {
	*CALData
	Attribute Attribute
}

// ICALDataRequestIdentify is the corresponding interface of CALDataRequestIdentify
type ICALDataRequestIdentify interface {
	ICALData
	// GetAttribute returns Attribute (property field)
	GetAttribute() Attribute
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *CALDataRequestIdentify) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

func (m *CALDataRequestIdentify) GetParent() *CALData {
	return m.CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *CALDataRequestIdentify) GetAttribute() Attribute {
	return m.Attribute
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataRequestIdentify factory function for CALDataRequestIdentify
func NewCALDataRequestIdentify(attribute Attribute, commandTypeContainer CALCommandTypeContainer) *CALDataRequestIdentify {
	_result := &CALDataRequestIdentify{
		Attribute: attribute,
		CALData:   NewCALData(commandTypeContainer),
	}
	_result.Child = _result
	return _result
}

func CastCALDataRequestIdentify(structType interface{}) *CALDataRequestIdentify {
	if casted, ok := structType.(CALDataRequestIdentify); ok {
		return &casted
	}
	if casted, ok := structType.(*CALDataRequestIdentify); ok {
		return casted
	}
	if casted, ok := structType.(CALData); ok {
		return CastCALDataRequestIdentify(casted.Child)
	}
	if casted, ok := structType.(*CALData); ok {
		return CastCALDataRequestIdentify(casted.Child)
	}
	return nil
}

func (m *CALDataRequestIdentify) GetTypeName() string {
	return "CALDataRequestIdentify"
}

func (m *CALDataRequestIdentify) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataRequestIdentify) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (attribute)
	lengthInBits += 8

	return lengthInBits
}

func (m *CALDataRequestIdentify) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataRequestIdentifyParse(readBuffer utils.ReadBuffer) (*CALDataRequestIdentify, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataRequestIdentify"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (attribute)
	if pullErr := readBuffer.PullContext("attribute"); pullErr != nil {
		return nil, pullErr
	}
	_attribute, _attributeErr := AttributeParse(readBuffer)
	if _attributeErr != nil {
		return nil, errors.Wrap(_attributeErr, "Error parsing 'attribute' field")
	}
	attribute := _attribute
	if closeErr := readBuffer.CloseContext("attribute"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CALDataRequestIdentify"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALDataRequestIdentify{
		Attribute: attribute,
		CALData:   &CALData{},
	}
	_child.CALData.Child = _child
	return _child, nil
}

func (m *CALDataRequestIdentify) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRequestIdentify"); pushErr != nil {
			return pushErr
		}

		// Simple Field (attribute)
		if pushErr := writeBuffer.PushContext("attribute"); pushErr != nil {
			return pushErr
		}
		_attributeErr := m.Attribute.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("attribute"); popErr != nil {
			return popErr
		}
		if _attributeErr != nil {
			return errors.Wrap(_attributeErr, "Error serializing 'attribute' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRequestIdentify"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataRequestIdentify) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
