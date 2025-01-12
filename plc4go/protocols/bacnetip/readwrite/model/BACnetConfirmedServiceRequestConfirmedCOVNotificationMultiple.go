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

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the corresponding interface of BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier (property field)
	GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetTimeStampEnclosed
	// GetListOfCovNotifications returns ListOfCovNotifications (property field)
	GetListOfCovNotifications() ListOfCovNotificationsList
	// IsBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple()
	// CreateBuilder creates a BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder() BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
}

// _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  BACnetContextTagObjectIdentifier
	TimeRemaining               BACnetContextTagUnsignedInteger
	Timestamp                   BACnetTimeStampEnclosed
	ListOfCovNotifications      ListOfCovNotificationsList
}

var _ BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple = (*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)(nil)

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple factory function for _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, timestamp BACnetTimeStampEnclosed, listOfCovNotifications ListOfCovNotificationsList, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	if initiatingDeviceIdentifier == nil {
		panic("initiatingDeviceIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	if timeRemaining == nil {
		panic("timeRemaining of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	if listOfCovNotifications == nil {
		panic("listOfCovNotifications of type ListOfCovNotificationsList for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:            initiatingDeviceIdentifier,
		TimeRemaining:                         timeRemaining,
		Timestamp:                             timestamp,
		ListOfCovNotifications:                listOfCovNotifications,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder is a builder for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
type BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, listOfCovNotifications ListOfCovNotificationsList) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithSubscriberProcessIdentifier adds SubscriberProcessIdentifier (property field)
	WithSubscriberProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithSubscriberProcessIdentifierBuilder adds SubscriberProcessIdentifier (property field) which is build by the builder
	WithSubscriberProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithInitiatingDeviceIdentifier adds InitiatingDeviceIdentifier (property field)
	WithInitiatingDeviceIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithInitiatingDeviceIdentifierBuilder adds InitiatingDeviceIdentifier (property field) which is build by the builder
	WithInitiatingDeviceIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithTimeRemaining adds TimeRemaining (property field)
	WithTimeRemaining(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithTimeRemainingBuilder adds TimeRemaining (property field) which is build by the builder
	WithTimeRemainingBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithTimestamp adds Timestamp (property field)
	WithOptionalTimestamp(BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithOptionalTimestampBuilder adds Timestamp (property field) which is build by the builder
	WithOptionalTimestampBuilder(func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithListOfCovNotifications adds ListOfCovNotifications (property field)
	WithListOfCovNotifications(ListOfCovNotificationsList) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// WithListOfCovNotificationsBuilder adds ListOfCovNotifications (property field) which is build by the builder
	WithListOfCovNotificationsBuilder(func(ListOfCovNotificationsListBuilder) ListOfCovNotificationsListBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConfirmedServiceRequestBuilder
	// Build builds the BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
}

// NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder() creates a BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
func NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder() BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	return &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder{_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple: new(_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple)}
}

type _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder struct {
	*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) = (*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
	contract.(*_BACnetConfirmedServiceRequest)._SubType = b._BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier BACnetContextTagObjectIdentifier, timeRemaining BACnetContextTagUnsignedInteger, listOfCovNotifications ListOfCovNotificationsList) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	return b.WithSubscriberProcessIdentifier(subscriberProcessIdentifier).WithInitiatingDeviceIdentifier(initiatingDeviceIdentifier).WithTimeRemaining(timeRemaining).WithListOfCovNotifications(listOfCovNotifications)
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithSubscriberProcessIdentifier(subscriberProcessIdentifier BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	b.SubscriberProcessIdentifier = subscriberProcessIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithSubscriberProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.SubscriberProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.SubscriberProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithInitiatingDeviceIdentifier(initiatingDeviceIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	b.InitiatingDeviceIdentifier = initiatingDeviceIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithInitiatingDeviceIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.InitiatingDeviceIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.InitiatingDeviceIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithTimeRemaining(timeRemaining BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	b.TimeRemaining = timeRemaining
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithTimeRemainingBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.TimeRemaining.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeRemaining, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithOptionalTimestamp(timestamp BACnetTimeStampEnclosed) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	b.Timestamp = timestamp
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithOptionalTimestampBuilder(builderSupplier func(BACnetTimeStampEnclosedBuilder) BACnetTimeStampEnclosedBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.Timestamp.CreateBACnetTimeStampEnclosedBuilder())
	var err error
	b.Timestamp, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTimeStampEnclosedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithListOfCovNotifications(listOfCovNotifications ListOfCovNotificationsList) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	b.ListOfCovNotifications = listOfCovNotifications
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) WithListOfCovNotificationsBuilder(builderSupplier func(ListOfCovNotificationsListBuilder) ListOfCovNotificationsListBuilder) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	builder := builderSupplier(b.ListOfCovNotifications.CreateListOfCovNotificationsListBuilder())
	var err error
	b.ListOfCovNotifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ListOfCovNotificationsListBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) Build() (BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, error) {
	if b.SubscriberProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'subscriberProcessIdentifier' not set"))
	}
	if b.InitiatingDeviceIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'initiatingDeviceIdentifier' not set"))
	}
	if b.TimeRemaining == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeRemaining' not set"))
	}
	if b.ListOfCovNotifications == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfCovNotifications' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) MustBuild() BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConfirmedServiceRequestBuilder().(*_BACnetConfirmedServiceRequestBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder().(*_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder creates a BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder
func (b *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) CreateBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder() BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder()
	}
	return &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleBuilder{_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_CONFIRMED_COV_NOTIFICATION_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetInitiatingDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTimestamp() BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetListOfCovNotifications() ListOfCovNotificationsList {
	return m.ListOfCovNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple(structType any) BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits(ctx)

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits(ctx)

	// Optional Field (timestamp)
	if m.Timestamp != nil {
		lengthInBits += m.Timestamp.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovNotifications)
	lengthInBits += m.ListOfCovNotifications.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	initiatingDeviceIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initiatingDeviceIdentifier' field"))
	}
	m.InitiatingDeviceIdentifier = initiatingDeviceIdentifier

	timeRemaining, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeRemaining' field"))
	}
	m.TimeRemaining = timeRemaining

	var timestamp BACnetTimeStampEnclosed
	_timestamp, err := ReadOptionalField[BACnetTimeStampEnclosed](ctx, "timestamp", ReadComplex[BACnetTimeStampEnclosed](BACnetTimeStampEnclosedParseWithBufferProducer((uint8)(uint8(3))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	if _timestamp != nil {
		timestamp = *_timestamp
		m.Timestamp = timestamp
	}

	listOfCovNotifications, err := ReadSimpleField[ListOfCovNotificationsList](ctx, "listOfCovNotifications", ReadComplex[ListOfCovNotificationsList](ListOfCovNotificationsListParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovNotifications' field"))
	}
	m.ListOfCovNotifications = listOfCovNotifications

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "initiatingDeviceIdentifier", m.GetInitiatingDeviceIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeRemaining", m.GetTimeRemaining(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeRemaining' field")
		}

		if err := WriteOptionalField[BACnetTimeStampEnclosed](ctx, "timestamp", GetRef(m.GetTimestamp()), WriteComplex[BACnetTimeStampEnclosed](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[ListOfCovNotificationsList](ctx, "listOfCovNotifications", m.GetListOfCovNotifications(), WriteComplex[ListOfCovNotificationsList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfCovNotifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) IsBACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple() {
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) deepCopy() *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleCopy := &_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.SubscriberProcessIdentifier),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.InitiatingDeviceIdentifier),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeRemaining),
		utils.DeepCopy[BACnetTimeStampEnclosed](m.Timestamp),
		utils.DeepCopy[ListOfCovNotificationsList](m.ListOfCovNotifications),
	}
	_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleCopy.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestConfirmedCOVNotificationMultipleCopy
}

func (m *_BACnetConfirmedServiceRequestConfirmedCOVNotificationMultiple) String() string {
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
