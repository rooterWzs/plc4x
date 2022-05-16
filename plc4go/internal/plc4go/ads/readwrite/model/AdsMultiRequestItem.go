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

// AdsMultiRequestItem is the data-structure of this message
type AdsMultiRequestItem struct {
	Child IAdsMultiRequestItemChild
}

// IAdsMultiRequestItem is the corresponding interface of AdsMultiRequestItem
type IAdsMultiRequestItem interface {
	// GetIndexGroup returns IndexGroup (discriminator field)
	GetIndexGroup() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IAdsMultiRequestItemParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAdsMultiRequestItemChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *AdsMultiRequestItem)
	GetParent() *AdsMultiRequestItem

	GetTypeName() string
	IAdsMultiRequestItem
}

// NewAdsMultiRequestItem factory function for AdsMultiRequestItem
func NewAdsMultiRequestItem() *AdsMultiRequestItem {
	return &AdsMultiRequestItem{}
}

func CastAdsMultiRequestItem(structType interface{}) *AdsMultiRequestItem {
	if casted, ok := structType.(AdsMultiRequestItem); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsMultiRequestItem); ok {
		return casted
	}
	if casted, ok := structType.(IAdsMultiRequestItemChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *AdsMultiRequestItem) GetTypeName() string {
	return "AdsMultiRequestItem"
}

func (m *AdsMultiRequestItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsMultiRequestItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *AdsMultiRequestItem) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *AdsMultiRequestItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsMultiRequestItemParse(readBuffer utils.ReadBuffer, indexGroup uint32) (*AdsMultiRequestItem, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsMultiRequestItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type AdsMultiRequestItemChild interface {
		InitializeParent(*AdsMultiRequestItem)
		GetParent() *AdsMultiRequestItem
	}
	var _child AdsMultiRequestItemChild
	var typeSwitchError error
	switch {
	case indexGroup == uint32(61568): // AdsMultiRequestItemRead
		_child, typeSwitchError = AdsMultiRequestItemReadParse(readBuffer, indexGroup)
	case indexGroup == uint32(61569): // AdsMultiRequestItemWrite
		_child, typeSwitchError = AdsMultiRequestItemWriteParse(readBuffer, indexGroup)
	case indexGroup == uint32(61570): // AdsMultiRequestItemReadWrite
		_child, typeSwitchError = AdsMultiRequestItemReadWriteParse(readBuffer, indexGroup)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("AdsMultiRequestItem"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent())
	return _child.GetParent(), nil
}

func (m *AdsMultiRequestItem) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *AdsMultiRequestItem) SerializeParent(writeBuffer utils.WriteBuffer, child IAdsMultiRequestItem, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsMultiRequestItem"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsMultiRequestItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsMultiRequestItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
