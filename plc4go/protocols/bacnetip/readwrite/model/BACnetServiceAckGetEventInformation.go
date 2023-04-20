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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetServiceAckGetEventInformation is the corresponding interface of BACnetServiceAckGetEventInformation
type BACnetServiceAckGetEventInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetListOfEventSummaries returns ListOfEventSummaries (property field)
	GetListOfEventSummaries() BACnetEventSummariesList
	// GetMoreEvents returns MoreEvents (property field)
	GetMoreEvents() BACnetContextTagBoolean
}

// BACnetServiceAckGetEventInformationExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckGetEventInformation.
// This is useful for switch cases.
type BACnetServiceAckGetEventInformationExactly interface {
	BACnetServiceAckGetEventInformation
	isBACnetServiceAckGetEventInformation() bool
}

// _BACnetServiceAckGetEventInformation is the data-structure of this message
type _BACnetServiceAckGetEventInformation struct {
	*_BACnetServiceAck
	ListOfEventSummaries BACnetEventSummariesList
	MoreEvents           BACnetContextTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetEventInformation) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_EVENT_INFORMATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetEventInformation) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckGetEventInformation) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetEventInformation) GetListOfEventSummaries() BACnetEventSummariesList {
	return m.ListOfEventSummaries
}

func (m *_BACnetServiceAckGetEventInformation) GetMoreEvents() BACnetContextTagBoolean {
	return m.MoreEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetEventInformation factory function for _BACnetServiceAckGetEventInformation
func NewBACnetServiceAckGetEventInformation(listOfEventSummaries BACnetEventSummariesList, moreEvents BACnetContextTagBoolean, serviceAckLength uint32) *_BACnetServiceAckGetEventInformation {
	_result := &_BACnetServiceAckGetEventInformation{
		ListOfEventSummaries: listOfEventSummaries,
		MoreEvents:           moreEvents,
		_BACnetServiceAck:    NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetEventInformation(structType any) BACnetServiceAckGetEventInformation {
	if casted, ok := structType.(BACnetServiceAckGetEventInformation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetEventInformation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetEventInformation) GetTypeName() string {
	return "BACnetServiceAckGetEventInformation"
}

func (m *_BACnetServiceAckGetEventInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (listOfEventSummaries)
	lengthInBits += m.ListOfEventSummaries.GetLengthInBits(ctx)

	// Simple field (moreEvents)
	lengthInBits += m.MoreEvents.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetServiceAckGetEventInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckGetEventInformationParse(theBytes []byte, serviceAckLength uint32) (BACnetServiceAckGetEventInformation, error) {
	return BACnetServiceAckGetEventInformationParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckGetEventInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckGetEventInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetEventInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetEventInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (listOfEventSummaries)
	if pullErr := readBuffer.PullContext("listOfEventSummaries"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfEventSummaries")
	}
	_listOfEventSummaries, _listOfEventSummariesErr := BACnetEventSummariesListParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _listOfEventSummariesErr != nil {
		return nil, errors.Wrap(_listOfEventSummariesErr, "Error parsing 'listOfEventSummaries' field of BACnetServiceAckGetEventInformation")
	}
	listOfEventSummaries := _listOfEventSummaries.(BACnetEventSummariesList)
	if closeErr := readBuffer.CloseContext("listOfEventSummaries"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfEventSummaries")
	}

	// Simple Field (moreEvents)
	if pullErr := readBuffer.PullContext("moreEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for moreEvents")
	}
	_moreEvents, _moreEventsErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _moreEventsErr != nil {
		return nil, errors.Wrap(_moreEventsErr, "Error parsing 'moreEvents' field of BACnetServiceAckGetEventInformation")
	}
	moreEvents := _moreEvents.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("moreEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for moreEvents")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetEventInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetEventInformation")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckGetEventInformation{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ListOfEventSummaries: listOfEventSummaries,
		MoreEvents:           moreEvents,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckGetEventInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckGetEventInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetEventInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetEventInformation")
		}

		// Simple Field (listOfEventSummaries)
		if pushErr := writeBuffer.PushContext("listOfEventSummaries"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfEventSummaries")
		}
		_listOfEventSummariesErr := writeBuffer.WriteSerializable(ctx, m.GetListOfEventSummaries())
		if popErr := writeBuffer.PopContext("listOfEventSummaries"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfEventSummaries")
		}
		if _listOfEventSummariesErr != nil {
			return errors.Wrap(_listOfEventSummariesErr, "Error serializing 'listOfEventSummaries' field")
		}

		// Simple Field (moreEvents)
		if pushErr := writeBuffer.PushContext("moreEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for moreEvents")
		}
		_moreEventsErr := writeBuffer.WriteSerializable(ctx, m.GetMoreEvents())
		if popErr := writeBuffer.PopContext("moreEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for moreEvents")
		}
		if _moreEventsErr != nil {
			return errors.Wrap(_moreEventsErr, "Error serializing 'moreEvents' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetEventInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetEventInformation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetEventInformation) isBACnetServiceAckGetEventInformation() bool {
	return true
}

func (m *_BACnetServiceAckGetEventInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
