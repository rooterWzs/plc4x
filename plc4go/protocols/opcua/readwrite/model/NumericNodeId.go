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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// NumericNodeId is the corresponding interface of NumericNodeId
type NumericNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint16
	// GetIdentifier returns Identifier (property field)
	GetIdentifier() uint32
}

// NumericNodeIdExactly can be used when we want exactly this type and not a type which fulfills NumericNodeId.
// This is useful for switch cases.
type NumericNodeIdExactly interface {
	NumericNodeId
	isNumericNodeId() bool
}

// _NumericNodeId is the data-structure of this message
type _NumericNodeId struct {
	NamespaceIndex uint16
	Identifier     uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NumericNodeId) GetNamespaceIndex() uint16 {
	return m.NamespaceIndex
}

func (m *_NumericNodeId) GetIdentifier() uint32 {
	return m.Identifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNumericNodeId factory function for _NumericNodeId
func NewNumericNodeId(namespaceIndex uint16, identifier uint32) *_NumericNodeId {
	return &_NumericNodeId{NamespaceIndex: namespaceIndex, Identifier: identifier}
}

// Deprecated: use the interface for direct cast
func CastNumericNodeId(structType any) NumericNodeId {
	if casted, ok := structType.(NumericNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*NumericNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_NumericNodeId) GetTypeName() string {
	return "NumericNodeId"
}

func (m *_NumericNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (namespaceIndex)
	lengthInBits += 16

	// Simple field (identifier)
	lengthInBits += 32

	return lengthInBits
}

func (m *_NumericNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NumericNodeIdParse(ctx context.Context, theBytes []byte) (NumericNodeId, error) {
	return NumericNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NumericNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NumericNodeId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NumericNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NumericNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint16("namespaceIndex", 16)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of NumericNodeId")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (identifier)
	_identifier, _identifierErr := readBuffer.ReadUint32("identifier", 32)
	if _identifierErr != nil {
		return nil, errors.Wrap(_identifierErr, "Error parsing 'identifier' field of NumericNodeId")
	}
	identifier := _identifier

	if closeErr := readBuffer.CloseContext("NumericNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NumericNodeId")
	}

	// Create the instance
	return &_NumericNodeId{
		NamespaceIndex: namespaceIndex,
		Identifier:     identifier,
	}, nil
}

func (m *_NumericNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NumericNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("NumericNodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NumericNodeId")
	}

	// Simple Field (namespaceIndex)
	namespaceIndex := uint16(m.GetNamespaceIndex())
	_namespaceIndexErr := writeBuffer.WriteUint16("namespaceIndex", 16, (namespaceIndex))
	if _namespaceIndexErr != nil {
		return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
	}

	// Simple Field (identifier)
	identifier := uint32(m.GetIdentifier())
	_identifierErr := writeBuffer.WriteUint32("identifier", 32, (identifier))
	if _identifierErr != nil {
		return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
	}

	if popErr := writeBuffer.PopContext("NumericNodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NumericNodeId")
	}
	return nil
}

func (m *_NumericNodeId) isNumericNodeId() bool {
	return true
}

func (m *_NumericNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}