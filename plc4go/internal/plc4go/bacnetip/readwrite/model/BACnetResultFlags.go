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

// BACnetResultFlags is the data-structure of this message
type BACnetResultFlags struct {
	RawBits *BACnetContextTagBitString

	// Arguments.
	TagNumber uint8
}

// IBACnetResultFlags is the corresponding interface of BACnetResultFlags
type IBACnetResultFlags interface {
	// GetRawBits returns RawBits (property field)
	GetRawBits() *BACnetContextTagBitString
	// GetFirstItem returns FirstItem (virtual field)
	GetFirstItem() bool
	// GetLastItem returns LastItem (virtual field)
	GetLastItem() bool
	// GetMoreItems returns MoreItems (virtual field)
	GetMoreItems() bool
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

func (m *BACnetResultFlags) GetRawBits() *BACnetContextTagBitString {
	return m.RawBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetResultFlags) GetFirstItem() bool {
	return bool(m.GetRawBits().GetPayload().GetData()[0])
}

func (m *BACnetResultFlags) GetLastItem() bool {
	return bool(m.GetRawBits().GetPayload().GetData()[1])
}

func (m *BACnetResultFlags) GetMoreItems() bool {
	return bool(m.GetRawBits().GetPayload().GetData()[2])
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetResultFlags factory function for BACnetResultFlags
func NewBACnetResultFlags(rawBits *BACnetContextTagBitString, tagNumber uint8) *BACnetResultFlags {
	return &BACnetResultFlags{RawBits: rawBits, TagNumber: tagNumber}
}

func CastBACnetResultFlags(structType interface{}) *BACnetResultFlags {
	if casted, ok := structType.(BACnetResultFlags); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetResultFlags); ok {
		return casted
	}
	return nil
}

func (m *BACnetResultFlags) GetTypeName() string {
	return "BACnetResultFlags"
}

func (m *BACnetResultFlags) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetResultFlags) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawBits)
	lengthInBits += m.RawBits.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetResultFlags) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetResultFlagsParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetResultFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetResultFlags"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (rawBits)
	if pullErr := readBuffer.PullContext("rawBits"); pullErr != nil {
		return nil, pullErr
	}
	_rawBits, _rawBitsErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_BIT_STRING))
	if _rawBitsErr != nil {
		return nil, errors.Wrap(_rawBitsErr, "Error parsing 'rawBits' field")
	}
	rawBits := CastBACnetContextTagBitString(_rawBits)
	if closeErr := readBuffer.CloseContext("rawBits"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_firstItem := rawBits.GetPayload().GetData()[0]
	firstItem := bool(_firstItem)
	_ = firstItem

	// Virtual field
	_lastItem := rawBits.GetPayload().GetData()[1]
	lastItem := bool(_lastItem)
	_ = lastItem

	// Virtual field
	_moreItems := rawBits.GetPayload().GetData()[2]
	moreItems := bool(_moreItems)
	_ = moreItems

	if closeErr := readBuffer.CloseContext("BACnetResultFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetResultFlags(rawBits, tagNumber), nil
}

func (m *BACnetResultFlags) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetResultFlags"); pushErr != nil {
		return pushErr
	}

	// Simple Field (rawBits)
	if pushErr := writeBuffer.PushContext("rawBits"); pushErr != nil {
		return pushErr
	}
	_rawBitsErr := m.RawBits.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("rawBits"); popErr != nil {
		return popErr
	}
	if _rawBitsErr != nil {
		return errors.Wrap(_rawBitsErr, "Error serializing 'rawBits' field")
	}
	// Virtual field
	if _firstItemErr := writeBuffer.WriteVirtual("firstItem", m.GetFirstItem()); _firstItemErr != nil {
		return errors.Wrap(_firstItemErr, "Error serializing 'firstItem' field")
	}
	// Virtual field
	if _lastItemErr := writeBuffer.WriteVirtual("lastItem", m.GetLastItem()); _lastItemErr != nil {
		return errors.Wrap(_lastItemErr, "Error serializing 'lastItem' field")
	}
	// Virtual field
	if _moreItemsErr := writeBuffer.WriteVirtual("moreItems", m.GetMoreItems()); _moreItemsErr != nil {
		return errors.Wrap(_moreItemsErr, "Error serializing 'moreItems' field")
	}

	if popErr := writeBuffer.PopContext("BACnetResultFlags"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetResultFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
