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

// BACnetFaultParameterFaultExtendedParametersEntryBitString is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryBitString
type BACnetFaultParameterFaultExtendedParametersEntryBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
	// IsBACnetFaultParameterFaultExtendedParametersEntryBitString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultExtendedParametersEntryBitString()
	// CreateBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
	CreateBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
}

// _BACnetFaultParameterFaultExtendedParametersEntryBitString is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryBitString struct {
	BACnetFaultParameterFaultExtendedParametersEntryContract
	BitStringValue BACnetApplicationTagBitString
}

var _ BACnetFaultParameterFaultExtendedParametersEntryBitString = (*_BACnetFaultParameterFaultExtendedParametersEntryBitString)(nil)
var _ BACnetFaultParameterFaultExtendedParametersEntryRequirements = (*_BACnetFaultParameterFaultExtendedParametersEntryBitString)(nil)

// NewBACnetFaultParameterFaultExtendedParametersEntryBitString factory function for _BACnetFaultParameterFaultExtendedParametersEntryBitString
func NewBACnetFaultParameterFaultExtendedParametersEntryBitString(peekedTagHeader BACnetTagHeader, bitStringValue BACnetApplicationTagBitString) *_BACnetFaultParameterFaultExtendedParametersEntryBitString {
	if bitStringValue == nil {
		panic("bitStringValue of type BACnetApplicationTagBitString for BACnetFaultParameterFaultExtendedParametersEntryBitString must not be nil")
	}
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryBitString{
		BACnetFaultParameterFaultExtendedParametersEntryContract: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
		BitStringValue: bitStringValue,
	}
	_result.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder is a builder for BACnetFaultParameterFaultExtendedParametersEntryBitString
type BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(bitStringValue BACnetApplicationTagBitString) BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
	// WithBitStringValue adds BitStringValue (property field)
	WithBitStringValue(BACnetApplicationTagBitString) BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
	// WithBitStringValueBuilder adds BitStringValue (property field) which is build by the builder
	WithBitStringValueBuilder(func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder
	// Build builds the BACnetFaultParameterFaultExtendedParametersEntryBitString or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultExtendedParametersEntryBitString, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultExtendedParametersEntryBitString
}

// NewBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder() creates a BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
func NewBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder {
	return &_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder{_BACnetFaultParameterFaultExtendedParametersEntryBitString: new(_BACnetFaultParameterFaultExtendedParametersEntryBitString)}
}

type _BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder struct {
	*_BACnetFaultParameterFaultExtendedParametersEntryBitString

	parentBuilder *_BACnetFaultParameterFaultExtendedParametersEntryBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) = (*_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder)(nil)

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) setParent(contract BACnetFaultParameterFaultExtendedParametersEntryContract) {
	b.BACnetFaultParameterFaultExtendedParametersEntryContract = contract
	contract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = b._BACnetFaultParameterFaultExtendedParametersEntryBitString
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) WithMandatoryFields(bitStringValue BACnetApplicationTagBitString) BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder {
	return b.WithBitStringValue(bitStringValue)
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) WithBitStringValue(bitStringValue BACnetApplicationTagBitString) BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder {
	b.BitStringValue = bitStringValue
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) WithBitStringValueBuilder(builderSupplier func(BACnetApplicationTagBitStringBuilder) BACnetApplicationTagBitStringBuilder) BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder {
	builder := builderSupplier(b.BitStringValue.CreateBACnetApplicationTagBitStringBuilder())
	var err error
	b.BitStringValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBitStringBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) Build() (BACnetFaultParameterFaultExtendedParametersEntryBitString, error) {
	if b.BitStringValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'bitStringValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultExtendedParametersEntryBitString.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) MustBuild() BACnetFaultParameterFaultExtendedParametersEntryBitString {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) Done() BACnetFaultParameterFaultExtendedParametersEntryBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetFaultParameterFaultExtendedParametersEntryBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) buildForBACnetFaultParameterFaultExtendedParametersEntry() (BACnetFaultParameterFaultExtendedParametersEntry, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder().(*_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder creates a BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder
func (b *_BACnetFaultParameterFaultExtendedParametersEntryBitString) CreateBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder() BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder()
	}
	return &_BACnetFaultParameterFaultExtendedParametersEntryBitStringBuilder{_BACnetFaultParameterFaultExtendedParametersEntryBitString: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetParent() BACnetFaultParameterFaultExtendedParametersEntryContract {
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryBitString(structType any) BACnetFaultParameterFaultExtendedParametersEntryBitString {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryBitString"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).getLengthInBits(ctx))

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultExtendedParametersEntry) (__bACnetFaultParameterFaultExtendedParametersEntryBitString BACnetFaultParameterFaultExtendedParametersEntryBitString, err error) {
	m.BACnetFaultParameterFaultExtendedParametersEntryContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	bitStringValue, err := ReadSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", ReadComplex[BACnetApplicationTagBitString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBitString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bitStringValue' field"))
	}
	m.BitStringValue = bitStringValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryBitString")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryBitString")
		}

		if err := WriteSimpleField[BACnetApplicationTagBitString](ctx, "bitStringValue", m.GetBitStringValue(), WriteComplex[BACnetApplicationTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryBitString")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) IsBACnetFaultParameterFaultExtendedParametersEntryBitString() {
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) deepCopy() *_BACnetFaultParameterFaultExtendedParametersEntryBitString {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultExtendedParametersEntryBitStringCopy := &_BACnetFaultParameterFaultExtendedParametersEntryBitString{
		m.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBitString](m.BitStringValue),
	}
	_BACnetFaultParameterFaultExtendedParametersEntryBitStringCopy.BACnetFaultParameterFaultExtendedParametersEntryContract.(*_BACnetFaultParameterFaultExtendedParametersEntry)._SubType = m
	return _BACnetFaultParameterFaultExtendedParametersEntryBitStringCopy
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) String() string {
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
