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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// StatusRequestLevel is the data-structure of this message
type StatusRequestLevel struct {
	*StatusRequest
	Application               byte
	StartingGroupAddressLabel byte
}

// IStatusRequestLevel is the corresponding interface of StatusRequestLevel
type IStatusRequestLevel interface {
	IStatusRequest
	// GetApplication returns Application (property field)
	GetApplication() byte
	// GetStartingGroupAddressLabel returns StartingGroupAddressLabel (property field)
	GetStartingGroupAddressLabel() byte
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

func (m *StatusRequestLevel) InitializeParent(parent *StatusRequest, statusType byte) {
	m.StatusRequest.StatusType = statusType
}

func (m *StatusRequestLevel) GetParent() *StatusRequest {
	return m.StatusRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *StatusRequestLevel) GetApplication() byte {
	return m.Application
}

func (m *StatusRequestLevel) GetStartingGroupAddressLabel() byte {
	return m.StartingGroupAddressLabel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusRequestLevel factory function for StatusRequestLevel
func NewStatusRequestLevel(application byte, startingGroupAddressLabel byte, statusType byte) *StatusRequestLevel {
	_result := &StatusRequestLevel{
		Application:               application,
		StartingGroupAddressLabel: startingGroupAddressLabel,
		StatusRequest:             NewStatusRequest(statusType),
	}
	_result.Child = _result
	return _result
}

func CastStatusRequestLevel(structType interface{}) *StatusRequestLevel {
	if casted, ok := structType.(StatusRequestLevel); ok {
		return &casted
	}
	if casted, ok := structType.(*StatusRequestLevel); ok {
		return casted
	}
	if casted, ok := structType.(StatusRequest); ok {
		return CastStatusRequestLevel(casted.Child)
	}
	if casted, ok := structType.(*StatusRequest); ok {
		return CastStatusRequestLevel(casted.Child)
	}
	return nil
}

func (m *StatusRequestLevel) GetTypeName() string {
	return "StatusRequestLevel"
}

func (m *StatusRequestLevel) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *StatusRequestLevel) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (startingGroupAddressLabel)
	lengthInBits += 8

	return lengthInBits
}

func (m *StatusRequestLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusRequestLevelParse(readBuffer utils.ReadBuffer) (*StatusRequestLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequestLevel"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x73) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x73),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0x07) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x07),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (application)
	_application, _applicationErr := readBuffer.ReadByte("application")
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application

	// Simple Field (startingGroupAddressLabel)
	_startingGroupAddressLabel, _startingGroupAddressLabelErr := readBuffer.ReadByte("startingGroupAddressLabel")
	if _startingGroupAddressLabelErr != nil {
		return nil, errors.Wrap(_startingGroupAddressLabelErr, "Error parsing 'startingGroupAddressLabel' field")
	}
	startingGroupAddressLabel := _startingGroupAddressLabel

	// Validation
	if !(bool(bool(bool(bool(bool(bool(bool(bool((startingGroupAddressLabel) == (0x00))) || bool(bool((startingGroupAddressLabel) == (0x20)))) || bool(bool((startingGroupAddressLabel) == (0x40)))) || bool(bool((startingGroupAddressLabel) == (0x60)))) || bool(bool((startingGroupAddressLabel) == (0x80)))) || bool(bool((startingGroupAddressLabel) == (0xA0)))) || bool(bool((startingGroupAddressLabel) == (0xC0)))) || bool(bool((startingGroupAddressLabel) == (0xE0)))) {
		return nil, utils.ParseValidationError{"invalid label"}
	}

	if closeErr := readBuffer.CloseContext("StatusRequestLevel"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &StatusRequestLevel{
		Application:               application,
		StartingGroupAddressLabel: startingGroupAddressLabel,
		StatusRequest:             &StatusRequest{},
	}
	_child.StatusRequest.Child = _child
	return _child, nil
}

func (m *StatusRequestLevel) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusRequestLevel"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x73))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x07))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (application)
		application := byte(m.Application)
		_applicationErr := writeBuffer.WriteByte("application", (application))
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Simple Field (startingGroupAddressLabel)
		startingGroupAddressLabel := byte(m.StartingGroupAddressLabel)
		_startingGroupAddressLabelErr := writeBuffer.WriteByte("startingGroupAddressLabel", (startingGroupAddressLabel))
		if _startingGroupAddressLabelErr != nil {
			return errors.Wrap(_startingGroupAddressLabelErr, "Error serializing 'startingGroupAddressLabel' field")
		}

		if popErr := writeBuffer.PopContext("StatusRequestLevel"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *StatusRequestLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
