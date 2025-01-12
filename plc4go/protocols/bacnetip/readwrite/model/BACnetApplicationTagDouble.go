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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetApplicationTagDouble is the corresponding interface of BACnetApplicationTagDouble
type BACnetApplicationTagDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float64
	// IsBACnetApplicationTagDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagDouble()
	// CreateBuilder creates a BACnetApplicationTagDoubleBuilder
	CreateBACnetApplicationTagDoubleBuilder() BACnetApplicationTagDoubleBuilder
}

// _BACnetApplicationTagDouble is the data-structure of this message
type _BACnetApplicationTagDouble struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadDouble
}

var _ BACnetApplicationTagDouble = (*_BACnetApplicationTagDouble)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagDouble)(nil)

// NewBACnetApplicationTagDouble factory function for _BACnetApplicationTagDouble
func NewBACnetApplicationTagDouble(header BACnetTagHeader, payload BACnetTagPayloadDouble) *_BACnetApplicationTagDouble {
	if payload == nil {
		panic("payload of type BACnetTagPayloadDouble for BACnetApplicationTagDouble must not be nil")
	}
	_result := &_BACnetApplicationTagDouble{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetApplicationTagDoubleBuilder is a builder for BACnetApplicationTagDouble
type BACnetApplicationTagDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(payload BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder
	// WithPayload adds Payload (property field)
	WithPayload(BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder
	// WithPayloadBuilder adds Payload (property field) which is build by the builder
	WithPayloadBuilder(func(BACnetTagPayloadDoubleBuilder) BACnetTagPayloadDoubleBuilder) BACnetApplicationTagDoubleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetApplicationTagBuilder
	// Build builds the BACnetApplicationTagDouble or returns an error if something is wrong
	Build() (BACnetApplicationTagDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetApplicationTagDouble
}

// NewBACnetApplicationTagDoubleBuilder() creates a BACnetApplicationTagDoubleBuilder
func NewBACnetApplicationTagDoubleBuilder() BACnetApplicationTagDoubleBuilder {
	return &_BACnetApplicationTagDoubleBuilder{_BACnetApplicationTagDouble: new(_BACnetApplicationTagDouble)}
}

type _BACnetApplicationTagDoubleBuilder struct {
	*_BACnetApplicationTagDouble

	parentBuilder *_BACnetApplicationTagBuilder

	err *utils.MultiError
}

var _ (BACnetApplicationTagDoubleBuilder) = (*_BACnetApplicationTagDoubleBuilder)(nil)

func (b *_BACnetApplicationTagDoubleBuilder) setParent(contract BACnetApplicationTagContract) {
	b.BACnetApplicationTagContract = contract
	contract.(*_BACnetApplicationTag)._SubType = b._BACnetApplicationTagDouble
}

func (b *_BACnetApplicationTagDoubleBuilder) WithMandatoryFields(payload BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder {
	return b.WithPayload(payload)
}

func (b *_BACnetApplicationTagDoubleBuilder) WithPayload(payload BACnetTagPayloadDouble) BACnetApplicationTagDoubleBuilder {
	b.Payload = payload
	return b
}

func (b *_BACnetApplicationTagDoubleBuilder) WithPayloadBuilder(builderSupplier func(BACnetTagPayloadDoubleBuilder) BACnetTagPayloadDoubleBuilder) BACnetApplicationTagDoubleBuilder {
	builder := builderSupplier(b.Payload.CreateBACnetTagPayloadDoubleBuilder())
	var err error
	b.Payload, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagPayloadDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetApplicationTagDoubleBuilder) Build() (BACnetApplicationTagDouble, error) {
	if b.Payload == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'payload' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetApplicationTagDouble.deepCopy(), nil
}

func (b *_BACnetApplicationTagDoubleBuilder) MustBuild() BACnetApplicationTagDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetApplicationTagDoubleBuilder) Done() BACnetApplicationTagBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetApplicationTagBuilder().(*_BACnetApplicationTagBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetApplicationTagDoubleBuilder) buildForBACnetApplicationTag() (BACnetApplicationTag, error) {
	return b.Build()
}

func (b *_BACnetApplicationTagDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetApplicationTagDoubleBuilder().(*_BACnetApplicationTagDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetApplicationTagDoubleBuilder creates a BACnetApplicationTagDoubleBuilder
func (b *_BACnetApplicationTagDouble) CreateBACnetApplicationTagDoubleBuilder() BACnetApplicationTagDoubleBuilder {
	if b == nil {
		return NewBACnetApplicationTagDoubleBuilder()
	}
	return &_BACnetApplicationTagDoubleBuilder{_BACnetApplicationTagDouble: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagDouble) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagDouble) GetPayload() BACnetTagPayloadDouble {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagDouble) GetActualValue() float64 {
	ctx := context.Background()
	_ = ctx
	return float64(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagDouble(structType any) BACnetApplicationTagDouble {
	if casted, ok := structType.(BACnetApplicationTagDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagDouble) GetTypeName() string {
	return "BACnetApplicationTagDouble"
}

func (m *_BACnetApplicationTagDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag) (__bACnetApplicationTagDouble BACnetApplicationTagDouble, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadDouble](ctx, "payload", ReadComplex[BACnetTagPayloadDouble](BACnetTagPayloadDoubleParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	actualValue, err := ReadVirtualField[float64](ctx, "actualValue", (*float64)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagDouble")
	}

	return m, nil
}

func (m *_BACnetApplicationTagDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagDouble")
		}

		if err := WriteSimpleField[BACnetTagPayloadDouble](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagDouble")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagDouble) IsBACnetApplicationTagDouble() {}

func (m *_BACnetApplicationTagDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetApplicationTagDouble) deepCopy() *_BACnetApplicationTagDouble {
	if m == nil {
		return nil
	}
	_BACnetApplicationTagDoubleCopy := &_BACnetApplicationTagDouble{
		m.BACnetApplicationTagContract.(*_BACnetApplicationTag).deepCopy(),
		utils.DeepCopy[BACnetTagPayloadDouble](m.Payload),
	}
	_BACnetApplicationTagDoubleCopy.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = m
	return _BACnetApplicationTagDoubleCopy
}

func (m *_BACnetApplicationTagDouble) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
