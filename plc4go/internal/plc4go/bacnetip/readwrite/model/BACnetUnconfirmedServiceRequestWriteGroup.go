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

// BACnetUnconfirmedServiceRequestWriteGroup is the data-structure of this message
type BACnetUnconfirmedServiceRequestWriteGroup struct {
	*BACnetUnconfirmedServiceRequest

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetUnconfirmedServiceRequestWriteGroup is the corresponding interface of BACnetUnconfirmedServiceRequestWriteGroup
type IBACnetUnconfirmedServiceRequestWriteGroup interface {
	IBACnetUnconfirmedServiceRequest
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

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_WRITE_GROUP
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestWriteGroup) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

// NewBACnetUnconfirmedServiceRequestWriteGroup factory function for BACnetUnconfirmedServiceRequestWriteGroup
func NewBACnetUnconfirmedServiceRequestWriteGroup(serviceRequestLength uint16) *BACnetUnconfirmedServiceRequestWriteGroup {
	_result := &BACnetUnconfirmedServiceRequestWriteGroup{
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestWriteGroup(structType interface{}) *BACnetUnconfirmedServiceRequestWriteGroup {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestWriteGroup); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestWriteGroup); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestWriteGroup(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestWriteGroup(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWriteGroup"
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWriteGroupParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetUnconfirmedServiceRequestWriteGroup, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWriteGroup"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"TODO: implement me"}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWriteGroup"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWriteGroup{
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWriteGroup"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWriteGroup"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestWriteGroup) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
