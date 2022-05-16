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

// S7VarRequestParameterItem is the data-structure of this message
type S7VarRequestParameterItem struct {
	Child IS7VarRequestParameterItemChild
}

// IS7VarRequestParameterItem is the corresponding interface of S7VarRequestParameterItem
type IS7VarRequestParameterItem interface {
	// GetItemType returns ItemType (discriminator field)
	GetItemType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IS7VarRequestParameterItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IS7VarRequestParameterItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IS7VarRequestParameterItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *S7VarRequestParameterItem)
	GetParent() *S7VarRequestParameterItem

	GetTypeName() string
	IS7VarRequestParameterItem
}

// NewS7VarRequestParameterItem factory function for S7VarRequestParameterItem
func NewS7VarRequestParameterItem() *S7VarRequestParameterItem {
	return &S7VarRequestParameterItem{}
}

func CastS7VarRequestParameterItem(structType interface{}) *S7VarRequestParameterItem {
	if casted, ok := structType.(S7VarRequestParameterItem); ok {
		return &casted
	}
	if casted, ok := structType.(*S7VarRequestParameterItem); ok {
		return casted
	}
	if casted, ok := structType.(IS7VarRequestParameterItemChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *S7VarRequestParameterItem) GetTypeName() string {
	return "S7VarRequestParameterItem"
}

func (m *S7VarRequestParameterItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7VarRequestParameterItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *S7VarRequestParameterItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (itemType)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7VarRequestParameterItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7VarRequestParameterItemParse(readBuffer utils.ReadBuffer) (*S7VarRequestParameterItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7VarRequestParameterItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType, _itemTypeErr := readBuffer.ReadUint8("itemType", 8)
	if _itemTypeErr != nil {
		return nil, errors.Wrap(_itemTypeErr, "Error parsing 'itemType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7VarRequestParameterItemChild interface {
		InitializeParent(*S7VarRequestParameterItem)
		GetParent() *S7VarRequestParameterItem
	}
	var _child S7VarRequestParameterItemChild
	var typeSwitchError error
	switch {
	case itemType == 0x12: // S7VarRequestParameterItemAddress
		_child, typeSwitchError = S7VarRequestParameterItemAddressParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("S7VarRequestParameterItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *S7VarRequestParameterItem) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *S7VarRequestParameterItem) SerializeParent(writeBuffer utils.WriteBuffer, child IS7VarRequestParameterItem, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("S7VarRequestParameterItem"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (itemType) (Used as input to a switch field)
	itemType := uint8(child.GetItemType())
	_itemTypeErr := writeBuffer.WriteUint8("itemType", 8, (itemType))

	if _itemTypeErr != nil {
		return errors.Wrap(_itemTypeErr, "Error serializing 'itemType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7VarRequestParameterItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7VarRequestParameterItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
