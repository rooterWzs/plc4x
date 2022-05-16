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

// BACnetContextTagPropertyIdentifier is the data-structure of this message
type BACnetContextTagPropertyIdentifier struct {
	*BACnetContextTag
	PropertyIdentifier BACnetPropertyIdentifier
	ProprietaryValue   uint32

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
	ActualLength             uint32
}

// IBACnetContextTagPropertyIdentifier is the corresponding interface of BACnetContextTagPropertyIdentifier
type IBACnetContextTagPropertyIdentifier interface {
	IBACnetContextTag
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifier
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
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

func (m *BACnetContextTagPropertyIdentifier) GetDataType() BACnetDataType {
	return BACnetDataType_BACNET_PROPERTY_IDENTIFIER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetContextTagPropertyIdentifier) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

func (m *BACnetContextTagPropertyIdentifier) GetParent() *BACnetContextTag {
	return m.BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetContextTagPropertyIdentifier) GetPropertyIdentifier() BACnetPropertyIdentifier {
	return m.PropertyIdentifier
}

func (m *BACnetContextTagPropertyIdentifier) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetContextTagPropertyIdentifier) GetIsProprietary() bool {
	return bool(bool((m.GetPropertyIdentifier()) == (BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagPropertyIdentifier factory function for BACnetContextTagPropertyIdentifier
func NewBACnetContextTagPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifier, proprietaryValue uint32, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool, actualLength uint32) *BACnetContextTagPropertyIdentifier {
	_result := &BACnetContextTagPropertyIdentifier{
		PropertyIdentifier: propertyIdentifier,
		ProprietaryValue:   proprietaryValue,
		BACnetContextTag:   NewBACnetContextTag(header, tagNumberArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetContextTagPropertyIdentifier(structType interface{}) *BACnetContextTagPropertyIdentifier {
	if casted, ok := structType.(BACnetContextTagPropertyIdentifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagPropertyIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagPropertyIdentifier(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagPropertyIdentifier(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagPropertyIdentifier) GetTypeName() string {
	return "BACnetContextTagPropertyIdentifier"
}

func (m *BACnetContextTagPropertyIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagPropertyIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Manual Field (propertyIdentifier)
	lengthInBits += uint16(int32(m.GetActualLength()) * int32(int32(8)))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagPropertyIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagPropertyIdentifierParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, actualLength uint32) (*BACnetContextTagPropertyIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagPropertyIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Manual Field (propertyIdentifier)
	propertyIdentifier, _propertyIdentifierErr := ReadPropertyIdentifier(readBuffer, actualLength)
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field")
	}

	// Manual Field (proprietaryValue)
	proprietaryValue, _proprietaryValueErr := ReadProprietaryPropertyIdentifier(readBuffer, propertyIdentifier, actualLength)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}

	// Virtual field
	_isProprietary := bool((propertyIdentifier) == (BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	if closeErr := readBuffer.CloseContext("BACnetContextTagPropertyIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagPropertyIdentifier{
		PropertyIdentifier: propertyIdentifier,
		ProprietaryValue:   proprietaryValue,
		BACnetContextTag:   &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child, nil
}

func (m *BACnetContextTagPropertyIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagPropertyIdentifier"); pushErr != nil {
			return pushErr
		}

		// Manual Field (propertyIdentifier)
		_propertyIdentifierErr := WritePropertyIdentifier(writeBuffer, m.GetPropertyIdentifier())
		if _propertyIdentifierErr != nil {
			return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
		}

		// Manual Field (proprietaryValue)
		_proprietaryValueErr := WriteProprietaryPropertyIdentifier(writeBuffer, m.GetPropertyIdentifier(), m.GetProprietaryValue())
		if _proprietaryValueErr != nil {
			return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
		}
		// Virtual field
		if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
			return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagPropertyIdentifier"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagPropertyIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
