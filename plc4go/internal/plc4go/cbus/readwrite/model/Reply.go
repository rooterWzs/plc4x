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

// Reply is the data-structure of this message
type Reply struct {
	MagicByte byte
	Child     IReplyChild
}

// IReply is the corresponding interface of Reply
type IReply interface {
	// GetMagicByte returns MagicByte (property field)
	GetMagicByte() byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IReplyParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IReply, serializeChildFunction func() error) error
	GetTypeName() string
}

type IReplyChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *Reply, magicByte byte)
	GetParent() *Reply

	GetTypeName() string
	IReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *Reply) GetMagicByte() byte {
	return m.MagicByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewReply factory function for Reply
func NewReply(magicByte byte) *Reply {
	return &Reply{MagicByte: magicByte}
}

func CastReply(structType interface{}) *Reply {
	if casted, ok := structType.(Reply); ok {
		return &casted
	}
	if casted, ok := structType.(*Reply); ok {
		return casted
	}
	if casted, ok := structType.(IReplyChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *Reply) GetTypeName() string {
	return "Reply"
}

func (m *Reply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *Reply) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *Reply) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *Reply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ReplyParse(readBuffer utils.ReadBuffer) (*Reply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("Reply"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (magicByte)
	currentPos = positionAware.GetPos()
	magicByte, _err := readBuffer.ReadByte("magicByte")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'magicByte' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ReplyChild interface {
		InitializeParent(*Reply, byte)
		GetParent() *Reply
	}
	var _child ReplyChild
	var typeSwitchError error
	switch {
	case magicByte == 0x0: // CALReplyReply
		_child, typeSwitchError = CALReplyReplyParse(readBuffer)
	case magicByte == 0x0: // MonitoredSALReply
		_child, typeSwitchError = MonitoredSALReplyParse(readBuffer)
	case magicByte == 0x0: // ConfirmationReply
		_child, typeSwitchError = ConfirmationReplyParse(readBuffer)
	case magicByte == 0x0: // PowerUpReply
		_child, typeSwitchError = PowerUpReplyParse(readBuffer)
	case magicByte == 0x0: // ParameterChangeReply
		_child, typeSwitchError = ParameterChangeReplyParse(readBuffer)
	case magicByte == 0x0: // ExclamationMarkReply
		_child, typeSwitchError = ExclamationMarkReplyParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("Reply"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), magicByte)
	return _child.GetParent(), nil
}

func (m *Reply) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *Reply) SerializeParent(writeBuffer utils.WriteBuffer, child IReply, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("Reply"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("Reply"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *Reply) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
