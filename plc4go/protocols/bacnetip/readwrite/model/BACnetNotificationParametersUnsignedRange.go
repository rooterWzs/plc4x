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

// BACnetNotificationParametersUnsignedRange is the corresponding interface of BACnetNotificationParametersUnsignedRange
type BACnetNotificationParametersUnsignedRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() BACnetContextTagUnsignedInteger
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersUnsignedRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersUnsignedRange()
	// CreateBuilder creates a BACnetNotificationParametersUnsignedRangeBuilder
	CreateBACnetNotificationParametersUnsignedRangeBuilder() BACnetNotificationParametersUnsignedRangeBuilder
}

// _BACnetNotificationParametersUnsignedRange is the data-structure of this message
type _BACnetNotificationParametersUnsignedRange struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	SequenceNumber  BACnetContextTagUnsignedInteger
	StatusFlags     BACnetStatusFlagsTagged
	ExceededLimit   BACnetContextTagUnsignedInteger
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersUnsignedRange = (*_BACnetNotificationParametersUnsignedRange)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersUnsignedRange)(nil)

// NewBACnetNotificationParametersUnsignedRange factory function for _BACnetNotificationParametersUnsignedRange
func NewBACnetNotificationParametersUnsignedRange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, sequenceNumber BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersUnsignedRange {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersUnsignedRange must not be nil")
	}
	if sequenceNumber == nil {
		panic("sequenceNumber of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersUnsignedRange must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersUnsignedRange must not be nil")
	}
	if exceededLimit == nil {
		panic("exceededLimit of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersUnsignedRange must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersUnsignedRange must not be nil")
	}
	_result := &_BACnetNotificationParametersUnsignedRange{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		SequenceNumber:                       sequenceNumber,
		StatusFlags:                          statusFlags,
		ExceededLimit:                        exceededLimit,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersUnsignedRangeBuilder is a builder for BACnetNotificationParametersUnsignedRange
type BACnetNotificationParametersUnsignedRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, sequenceNumber BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersUnsignedRangeBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersUnsignedRangeBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersUnsignedRangeBuilder
	// WithSequenceNumber adds SequenceNumber (property field)
	WithSequenceNumber(BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedRangeBuilder
	// WithSequenceNumberBuilder adds SequenceNumber (property field) which is build by the builder
	WithSequenceNumberBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedRangeBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersUnsignedRangeBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersUnsignedRangeBuilder
	// WithExceededLimit adds ExceededLimit (property field)
	WithExceededLimit(BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedRangeBuilder
	// WithExceededLimitBuilder adds ExceededLimit (property field) which is build by the builder
	WithExceededLimitBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedRangeBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersUnsignedRangeBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersUnsignedRangeBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetNotificationParametersBuilder
	// Build builds the BACnetNotificationParametersUnsignedRange or returns an error if something is wrong
	Build() (BACnetNotificationParametersUnsignedRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersUnsignedRange
}

// NewBACnetNotificationParametersUnsignedRangeBuilder() creates a BACnetNotificationParametersUnsignedRangeBuilder
func NewBACnetNotificationParametersUnsignedRangeBuilder() BACnetNotificationParametersUnsignedRangeBuilder {
	return &_BACnetNotificationParametersUnsignedRangeBuilder{_BACnetNotificationParametersUnsignedRange: new(_BACnetNotificationParametersUnsignedRange)}
}

type _BACnetNotificationParametersUnsignedRangeBuilder struct {
	*_BACnetNotificationParametersUnsignedRange

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersUnsignedRangeBuilder) = (*_BACnetNotificationParametersUnsignedRangeBuilder)(nil)

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
	contract.(*_BACnetNotificationParameters)._SubType = b._BACnetNotificationParametersUnsignedRange
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, sequenceNumber BACnetContextTagUnsignedInteger, statusFlags BACnetStatusFlagsTagged, exceededLimit BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersUnsignedRangeBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithSequenceNumber(sequenceNumber).WithStatusFlags(statusFlags).WithExceededLimit(exceededLimit).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersUnsignedRangeBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersUnsignedRangeBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithSequenceNumber(sequenceNumber BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedRangeBuilder {
	b.SequenceNumber = sequenceNumber
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithSequenceNumberBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedRangeBuilder {
	builder := builderSupplier(b.SequenceNumber.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.SequenceNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersUnsignedRangeBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersUnsignedRangeBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithExceededLimit(exceededLimit BACnetContextTagUnsignedInteger) BACnetNotificationParametersUnsignedRangeBuilder {
	b.ExceededLimit = exceededLimit
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithExceededLimitBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersUnsignedRangeBuilder {
	builder := builderSupplier(b.ExceededLimit.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ExceededLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersUnsignedRangeBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersUnsignedRangeBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) Build() (BACnetNotificationParametersUnsignedRange, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.SequenceNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'sequenceNumber' not set"))
	}
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.ExceededLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'exceededLimit' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersUnsignedRange.deepCopy(), nil
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) MustBuild() BACnetNotificationParametersUnsignedRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) Done() BACnetNotificationParametersBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetNotificationParametersBuilder().(*_BACnetNotificationParametersBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersUnsignedRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersUnsignedRangeBuilder().(*_BACnetNotificationParametersUnsignedRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersUnsignedRangeBuilder creates a BACnetNotificationParametersUnsignedRangeBuilder
func (b *_BACnetNotificationParametersUnsignedRange) CreateBACnetNotificationParametersUnsignedRangeBuilder() BACnetNotificationParametersUnsignedRangeBuilder {
	if b == nil {
		return NewBACnetNotificationParametersUnsignedRangeBuilder()
	}
	return &_BACnetNotificationParametersUnsignedRangeBuilder{_BACnetNotificationParametersUnsignedRange: b.deepCopy()}
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

func (m *_BACnetNotificationParametersUnsignedRange) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersUnsignedRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersUnsignedRange) GetSequenceNumber() BACnetContextTagUnsignedInteger {
	return m.SequenceNumber
}

func (m *_BACnetNotificationParametersUnsignedRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersUnsignedRange) GetExceededLimit() BACnetContextTagUnsignedInteger {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersUnsignedRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersUnsignedRange(structType any) BACnetNotificationParametersUnsignedRange {
	if casted, ok := structType.(BACnetNotificationParametersUnsignedRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersUnsignedRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersUnsignedRange) GetTypeName() string {
	return "BACnetNotificationParametersUnsignedRange"
}

func (m *_BACnetNotificationParametersUnsignedRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (sequenceNumber)
	lengthInBits += m.SequenceNumber.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersUnsignedRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersUnsignedRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersUnsignedRange BACnetNotificationParametersUnsignedRange, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersUnsignedRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersUnsignedRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	sequenceNumber, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "sequenceNumber", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sequenceNumber' field"))
	}
	m.SequenceNumber = sequenceNumber

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	exceededLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "exceededLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceededLimit' field"))
	}
	m.ExceededLimit = exceededLimit

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersUnsignedRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersUnsignedRange")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersUnsignedRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersUnsignedRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersUnsignedRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersUnsignedRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "sequenceNumber", m.GetSequenceNumber(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sequenceNumber' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "exceededLimit", m.GetExceededLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceededLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersUnsignedRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersUnsignedRange")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersUnsignedRange) IsBACnetNotificationParametersUnsignedRange() {}

func (m *_BACnetNotificationParametersUnsignedRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersUnsignedRange) deepCopy() *_BACnetNotificationParametersUnsignedRange {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersUnsignedRangeCopy := &_BACnetNotificationParametersUnsignedRange{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.SequenceNumber),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.StatusFlags),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.ExceededLimit),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	_BACnetNotificationParametersUnsignedRangeCopy.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersUnsignedRangeCopy
}

func (m *_BACnetNotificationParametersUnsignedRange) String() string {
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
