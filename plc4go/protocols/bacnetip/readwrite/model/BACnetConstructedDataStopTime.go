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

// BACnetConstructedDataStopTime is the corresponding interface of BACnetConstructedDataStopTime
type BACnetConstructedDataStopTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStopTime returns StopTime (property field)
	GetStopTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataStopTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataStopTime.
// This is useful for switch cases.
type BACnetConstructedDataStopTimeExactly interface {
	BACnetConstructedDataStopTime
	isBACnetConstructedDataStopTime() bool
}

// _BACnetConstructedDataStopTime is the data-structure of this message
type _BACnetConstructedDataStopTime struct {
	*_BACnetConstructedData
	StopTime BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStopTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataStopTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStopTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataStopTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStopTime) GetStopTime() BACnetDateTime {
	return m.StopTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStopTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetStopTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStopTime factory function for _BACnetConstructedDataStopTime
func NewBACnetConstructedDataStopTime(stopTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataStopTime {
	_result := &_BACnetConstructedDataStopTime{
		StopTime:               stopTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStopTime(structType any) BACnetConstructedDataStopTime {
	if casted, ok := structType.(BACnetConstructedDataStopTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStopTime) GetTypeName() string {
	return "BACnetConstructedDataStopTime"
}

func (m *_BACnetConstructedDataStopTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (stopTime)
	lengthInBits += m.StopTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataStopTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataStopTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStopTime, error) {
	return BACnetConstructedDataStopTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataStopTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStopTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStopTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stopTime)
	if pullErr := readBuffer.PullContext("stopTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for stopTime")
	}
	_stopTime, _stopTimeErr := BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	if _stopTimeErr != nil {
		return nil, errors.Wrap(_stopTimeErr, "Error parsing 'stopTime' field of BACnetConstructedDataStopTime")
	}
	stopTime := _stopTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("stopTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for stopTime")
	}

	// Virtual field
	_actualValue := stopTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStopTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataStopTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		StopTime: stopTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataStopTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStopTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStopTime")
		}

		// Simple Field (stopTime)
		if pushErr := writeBuffer.PushContext("stopTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for stopTime")
		}
		_stopTimeErr := writeBuffer.WriteSerializable(ctx, m.GetStopTime())
		if popErr := writeBuffer.PopContext("stopTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for stopTime")
		}
		if _stopTimeErr != nil {
			return errors.Wrap(_stopTimeErr, "Error serializing 'stopTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStopTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataStopTime) isBACnetConstructedDataStopTime() bool {
	return true
}

func (m *_BACnetConstructedDataStopTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
