/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const RequestReset_TILDE byte = 0x7E

// RequestReset is the corresponding interface of RequestReset
type RequestReset interface {
	utils.LengthAware
	utils.Serializable
	Request
}

// RequestResetExactly can be used when we want exactly this type and not a type which fulfills RequestReset.
// This is useful for switch cases.
type RequestResetExactly interface {
	RequestReset
	isRequestReset() bool
}

// _RequestReset is the data-structure of this message
type _RequestReset struct {
	*_Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestReset) InitializeParent(parent Request, peekedByte RequestType, termination RequestTermination) {
	m.PeekedByte = peekedByte
	m.Termination = termination
}

func (m *_RequestReset) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestReset) GetTilde() byte {
	return RequestReset_TILDE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestReset factory function for _RequestReset
func NewRequestReset(peekedByte RequestType, termination RequestTermination, srchk bool, messageLength uint16) *_RequestReset {
	_result := &_RequestReset{
		_Request: NewRequest(peekedByte, termination, srchk, messageLength),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestReset(structType interface{}) RequestReset {
	if casted, ok := structType.(RequestReset); ok {
		return casted
	}
	if casted, ok := structType.(*RequestReset); ok {
		return *casted
	}
	return nil
}

func (m *_RequestReset) GetTypeName() string {
	return "RequestReset"
}

func (m *_RequestReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (tilde)
	lengthInBits += 8

	return lengthInBits
}

func (m *_RequestReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestResetParse(readBuffer utils.ReadBuffer, srchk bool, messageLength uint16) (RequestReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (tilde)
	tilde, _tildeErr := readBuffer.ReadByte("tilde")
	if _tildeErr != nil {
		return nil, errors.Wrap(_tildeErr, "Error parsing 'tilde' field of RequestReset")
	}
	if tilde != RequestReset_TILDE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestReset_TILDE) + " but got " + fmt.Sprintf("%d", tilde))
	}

	if closeErr := readBuffer.CloseContext("RequestReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestReset")
	}

	// Create a partially initialized instance
	_child := &_RequestReset{
		_Request: &_Request{
			Srchk:         srchk,
			MessageLength: messageLength,
		},
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestReset) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestReset")
		}

		// Const Field (tilde)
		_tildeErr := writeBuffer.WriteByte("tilde", 0x7E)
		if _tildeErr != nil {
			return errors.Wrap(_tildeErr, "Error serializing 'tilde' field")
		}

		if popErr := writeBuffer.PopContext("RequestReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestReset")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_RequestReset) isRequestReset() bool {
	return true
}

func (m *_RequestReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
