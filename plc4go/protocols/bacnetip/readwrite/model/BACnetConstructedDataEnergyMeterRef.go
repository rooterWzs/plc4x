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

// BACnetConstructedDataEnergyMeterRef is the corresponding interface of BACnetConstructedDataEnergyMeterRef
type BACnetConstructedDataEnergyMeterRef interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetEnergyMeterRef returns EnergyMeterRef (property field)
	GetEnergyMeterRef() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
	// IsBACnetConstructedDataEnergyMeterRef is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEnergyMeterRef()
	// CreateBuilder creates a BACnetConstructedDataEnergyMeterRefBuilder
	CreateBACnetConstructedDataEnergyMeterRefBuilder() BACnetConstructedDataEnergyMeterRefBuilder
}

// _BACnetConstructedDataEnergyMeterRef is the data-structure of this message
type _BACnetConstructedDataEnergyMeterRef struct {
	BACnetConstructedDataContract
	EnergyMeterRef BACnetDeviceObjectReference
}

var _ BACnetConstructedDataEnergyMeterRef = (*_BACnetConstructedDataEnergyMeterRef)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEnergyMeterRef)(nil)

// NewBACnetConstructedDataEnergyMeterRef factory function for _BACnetConstructedDataEnergyMeterRef
func NewBACnetConstructedDataEnergyMeterRef(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, energyMeterRef BACnetDeviceObjectReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEnergyMeterRef {
	if energyMeterRef == nil {
		panic("energyMeterRef of type BACnetDeviceObjectReference for BACnetConstructedDataEnergyMeterRef must not be nil")
	}
	_result := &_BACnetConstructedDataEnergyMeterRef{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		EnergyMeterRef:                energyMeterRef,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataEnergyMeterRefBuilder is a builder for BACnetConstructedDataEnergyMeterRef
type BACnetConstructedDataEnergyMeterRefBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(energyMeterRef BACnetDeviceObjectReference) BACnetConstructedDataEnergyMeterRefBuilder
	// WithEnergyMeterRef adds EnergyMeterRef (property field)
	WithEnergyMeterRef(BACnetDeviceObjectReference) BACnetConstructedDataEnergyMeterRefBuilder
	// WithEnergyMeterRefBuilder adds EnergyMeterRef (property field) which is build by the builder
	WithEnergyMeterRefBuilder(func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataEnergyMeterRefBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataEnergyMeterRef or returns an error if something is wrong
	Build() (BACnetConstructedDataEnergyMeterRef, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataEnergyMeterRef
}

// NewBACnetConstructedDataEnergyMeterRefBuilder() creates a BACnetConstructedDataEnergyMeterRefBuilder
func NewBACnetConstructedDataEnergyMeterRefBuilder() BACnetConstructedDataEnergyMeterRefBuilder {
	return &_BACnetConstructedDataEnergyMeterRefBuilder{_BACnetConstructedDataEnergyMeterRef: new(_BACnetConstructedDataEnergyMeterRef)}
}

type _BACnetConstructedDataEnergyMeterRefBuilder struct {
	*_BACnetConstructedDataEnergyMeterRef

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataEnergyMeterRefBuilder) = (*_BACnetConstructedDataEnergyMeterRefBuilder)(nil)

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataEnergyMeterRef
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) WithMandatoryFields(energyMeterRef BACnetDeviceObjectReference) BACnetConstructedDataEnergyMeterRefBuilder {
	return b.WithEnergyMeterRef(energyMeterRef)
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) WithEnergyMeterRef(energyMeterRef BACnetDeviceObjectReference) BACnetConstructedDataEnergyMeterRefBuilder {
	b.EnergyMeterRef = energyMeterRef
	return b
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) WithEnergyMeterRefBuilder(builderSupplier func(BACnetDeviceObjectReferenceBuilder) BACnetDeviceObjectReferenceBuilder) BACnetConstructedDataEnergyMeterRefBuilder {
	builder := builderSupplier(b.EnergyMeterRef.CreateBACnetDeviceObjectReferenceBuilder())
	var err error
	b.EnergyMeterRef, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetDeviceObjectReferenceBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) Build() (BACnetConstructedDataEnergyMeterRef, error) {
	if b.EnergyMeterRef == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'energyMeterRef' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataEnergyMeterRef.deepCopy(), nil
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) MustBuild() BACnetConstructedDataEnergyMeterRef {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataEnergyMeterRefBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataEnergyMeterRefBuilder().(*_BACnetConstructedDataEnergyMeterRefBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataEnergyMeterRefBuilder creates a BACnetConstructedDataEnergyMeterRefBuilder
func (b *_BACnetConstructedDataEnergyMeterRef) CreateBACnetConstructedDataEnergyMeterRefBuilder() BACnetConstructedDataEnergyMeterRefBuilder {
	if b == nil {
		return NewBACnetConstructedDataEnergyMeterRefBuilder()
	}
	return &_BACnetConstructedDataEnergyMeterRefBuilder{_BACnetConstructedDataEnergyMeterRef: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENERGY_METER_REF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetEnergyMeterRef() BACnetDeviceObjectReference {
	return m.EnergyMeterRef
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeterRef) GetActualValue() BACnetDeviceObjectReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectReference(m.GetEnergyMeterRef())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEnergyMeterRef(structType any) BACnetConstructedDataEnergyMeterRef {
	if casted, ok := structType.(BACnetConstructedDataEnergyMeterRef); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEnergyMeterRef); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetTypeName() string {
	return "BACnetConstructedDataEnergyMeterRef"
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (energyMeterRef)
	lengthInBits += m.EnergyMeterRef.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEnergyMeterRef) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEnergyMeterRef) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEnergyMeterRef BACnetConstructedDataEnergyMeterRef, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEnergyMeterRef"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEnergyMeterRef")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	energyMeterRef, err := ReadSimpleField[BACnetDeviceObjectReference](ctx, "energyMeterRef", ReadComplex[BACnetDeviceObjectReference](BACnetDeviceObjectReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'energyMeterRef' field"))
	}
	m.EnergyMeterRef = energyMeterRef

	actualValue, err := ReadVirtualField[BACnetDeviceObjectReference](ctx, "actualValue", (*BACnetDeviceObjectReference)(nil), energyMeterRef)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEnergyMeterRef"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEnergyMeterRef")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEnergyMeterRef) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEnergyMeterRef) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEnergyMeterRef"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEnergyMeterRef")
		}

		if err := WriteSimpleField[BACnetDeviceObjectReference](ctx, "energyMeterRef", m.GetEnergyMeterRef(), WriteComplex[BACnetDeviceObjectReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'energyMeterRef' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEnergyMeterRef"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEnergyMeterRef")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEnergyMeterRef) IsBACnetConstructedDataEnergyMeterRef() {}

func (m *_BACnetConstructedDataEnergyMeterRef) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataEnergyMeterRef) deepCopy() *_BACnetConstructedDataEnergyMeterRef {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataEnergyMeterRefCopy := &_BACnetConstructedDataEnergyMeterRef{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetDeviceObjectReference](m.EnergyMeterRef),
	}
	_BACnetConstructedDataEnergyMeterRefCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataEnergyMeterRefCopy
}

func (m *_BACnetConstructedDataEnergyMeterRef) String() string {
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
