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

// BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is the data-structure of this message
type BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber struct {
	*BACnetConfirmedServiceRequestReadRangeRange
	ReferenceSequenceNumber *BACnetApplicationTagUnsignedInteger
	Count                   *BACnetApplicationTagSignedInteger
}

// IBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
type IBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber interface {
	IBACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceSequenceNumber returns ReferenceSequenceNumber (property field)
	GetReferenceSequenceNumber() *BACnetApplicationTagUnsignedInteger
	// GetCount returns Count (property field)
	GetCount() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) InitializeParent(parent *BACnetConfirmedServiceRequestReadRangeRange, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConfirmedServiceRequestReadRangeRange.PeekedTagHeader = peekedTagHeader
	m.BACnetConfirmedServiceRequestReadRangeRange.OpeningTag = openingTag
	m.BACnetConfirmedServiceRequestReadRangeRange.ClosingTag = closingTag
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetParent() *BACnetConfirmedServiceRequestReadRangeRange {
	return m.BACnetConfirmedServiceRequestReadRangeRange
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetReferenceSequenceNumber() *BACnetApplicationTagUnsignedInteger {
	return m.ReferenceSequenceNumber
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetCount() *BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber factory function for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
func NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(referenceSequenceNumber *BACnetApplicationTagUnsignedInteger, count *BACnetApplicationTagSignedInteger, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	_result := &BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{
		ReferenceSequenceNumber: referenceSequenceNumber,
		Count:                   count,
		BACnetConfirmedServiceRequestReadRangeRange: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(structType interface{}) *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRange); ok {
		return CastBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRange); ok {
		return CastBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referenceSequenceNumber)
	lengthInBits += m.ReferenceSequenceNumber.GetLengthInBits()

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceSequenceNumber)
	if pullErr := readBuffer.PullContext("referenceSequenceNumber"); pullErr != nil {
		return nil, pullErr
	}
	_referenceSequenceNumber, _referenceSequenceNumberErr := BACnetApplicationTagParse(readBuffer)
	if _referenceSequenceNumberErr != nil {
		return nil, errors.Wrap(_referenceSequenceNumberErr, "Error parsing 'referenceSequenceNumber' field")
	}
	referenceSequenceNumber := CastBACnetApplicationTagUnsignedInteger(_referenceSequenceNumber)
	if closeErr := readBuffer.CloseContext("referenceSequenceNumber"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (count)
	if pullErr := readBuffer.PullContext("count"); pullErr != nil {
		return nil, pullErr
	}
	_count, _countErr := BACnetApplicationTagParse(readBuffer)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := CastBACnetApplicationTagSignedInteger(_count)
	if closeErr := readBuffer.CloseContext("count"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{
		ReferenceSequenceNumber: CastBACnetApplicationTagUnsignedInteger(referenceSequenceNumber),
		Count:                   CastBACnetApplicationTagSignedInteger(count),
		BACnetConfirmedServiceRequestReadRangeRange: &BACnetConfirmedServiceRequestReadRangeRange{},
	}
	_child.BACnetConfirmedServiceRequestReadRangeRange.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); pushErr != nil {
			return pushErr
		}

		// Simple Field (referenceSequenceNumber)
		if pushErr := writeBuffer.PushContext("referenceSequenceNumber"); pushErr != nil {
			return pushErr
		}
		_referenceSequenceNumberErr := m.ReferenceSequenceNumber.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("referenceSequenceNumber"); popErr != nil {
			return popErr
		}
		if _referenceSequenceNumberErr != nil {
			return errors.Wrap(_referenceSequenceNumberErr, "Error serializing 'referenceSequenceNumber' field")
		}

		// Simple Field (count)
		if pushErr := writeBuffer.PushContext("count"); pushErr != nil {
			return pushErr
		}
		_countErr := m.Count.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("count"); popErr != nil {
			return popErr
		}
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
