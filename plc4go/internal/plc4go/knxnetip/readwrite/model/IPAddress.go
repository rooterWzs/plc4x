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

// IPAddress is the data-structure of this message
type IPAddress struct {
	Addr []byte
}

// IIPAddress is the corresponding interface of IPAddress
type IIPAddress interface {
	// GetAddr returns Addr (property field)
	GetAddr() []byte
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

func (m *IPAddress) GetAddr() []byte {
	return m.Addr
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIPAddress factory function for IPAddress
func NewIPAddress(addr []byte) *IPAddress {
	return &IPAddress{Addr: addr}
}

func CastIPAddress(structType interface{}) *IPAddress {
	if casted, ok := structType.(IPAddress); ok {
		return &casted
	}
	if casted, ok := structType.(*IPAddress); ok {
		return casted
	}
	return nil
}

func (m *IPAddress) GetTypeName() string {
	return "IPAddress"
}

func (m *IPAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IPAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Addr) > 0 {
		lengthInBits += 8 * uint16(len(m.Addr))
	}

	return lengthInBits
}

func (m *IPAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IPAddressParse(readBuffer utils.ReadBuffer) (*IPAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IPAddress"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (addr)
	numberOfBytesaddr := int(uint16(4))
	addr, _readArrayErr := readBuffer.ReadByteArray("addr", numberOfBytesaddr)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'addr' field")
	}

	if closeErr := readBuffer.CloseContext("IPAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewIPAddress(addr), nil
}

func (m *IPAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("IPAddress"); pushErr != nil {
		return pushErr
	}

	// Array Field (addr)
	if m.Addr != nil {
		// Byte Array field (addr)
		_writeArrayErr := writeBuffer.WriteByteArray("addr", m.Addr)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'addr' field")
		}
	}

	if popErr := writeBuffer.PopContext("IPAddress"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *IPAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
