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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEventMessageTexts is the corresponding interface of BACnetConstructedDataEventMessageTexts
type BACnetConstructedDataEventMessageTexts interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetEventMessageTexts returns EventMessageTexts (property field)
	GetEventMessageTexts() []BACnetOptionalCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetToOffnormalText returns ToOffnormalText (virtual field)
	GetToOffnormalText() BACnetOptionalCharacterString
	// GetToFaultText returns ToFaultText (virtual field)
	GetToFaultText() BACnetOptionalCharacterString
	// GetToNormalText returns ToNormalText (virtual field)
	GetToNormalText() BACnetOptionalCharacterString
}

// BACnetConstructedDataEventMessageTextsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventMessageTexts.
// This is useful for switch cases.
type BACnetConstructedDataEventMessageTextsExactly interface {
	BACnetConstructedDataEventMessageTexts
	isBACnetConstructedDataEventMessageTexts() bool
}

// _BACnetConstructedDataEventMessageTexts is the data-structure of this message
type _BACnetConstructedDataEventMessageTexts struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	EventMessageTexts    []BACnetOptionalCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEventMessageTexts) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EVENT_MESSAGE_TEXTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventMessageTexts) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataEventMessageTexts) GetEventMessageTexts() []BACnetOptionalCharacterString {
	return m.EventMessageTexts
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEventMessageTexts) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

func (m *_BACnetConstructedDataEventMessageTexts) GetToOffnormalText() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[0]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *_BACnetConstructedDataEventMessageTexts) GetToFaultText() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[1]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

func (m *_BACnetConstructedDataEventMessageTexts) GetToNormalText() BACnetOptionalCharacterString {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return CastBACnetOptionalCharacterString(CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(m.GetEventMessageTexts())) == (3)), func() any { return CastBACnetOptionalCharacterString(m.GetEventMessageTexts()[2]) }, func() any { return CastBACnetOptionalCharacterString(nil) })))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventMessageTexts factory function for _BACnetConstructedDataEventMessageTexts
func NewBACnetConstructedDataEventMessageTexts(numberOfDataElements BACnetApplicationTagUnsignedInteger, eventMessageTexts []BACnetOptionalCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventMessageTexts {
	_result := &_BACnetConstructedDataEventMessageTexts{
		NumberOfDataElements:   numberOfDataElements,
		EventMessageTexts:      eventMessageTexts,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventMessageTexts(structType any) BACnetConstructedDataEventMessageTexts {
	if casted, ok := structType.(BACnetConstructedDataEventMessageTexts); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventMessageTexts); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventMessageTexts) GetTypeName() string {
	return "BACnetConstructedDataEventMessageTexts"
}

func (m *_BACnetConstructedDataEventMessageTexts) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.EventMessageTexts) > 0 {
		for _, element := range m.EventMessageTexts {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEventMessageTexts) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataEventMessageTextsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventMessageTexts, error) {
	return BACnetConstructedDataEventMessageTextsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataEventMessageTextsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventMessageTexts, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventMessageTexts"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventMessageTexts")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataEventMessageTexts")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (eventMessageTexts)
	if pullErr := readBuffer.PullContext("eventMessageTexts", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventMessageTexts")
	}
	// Terminated array
	var eventMessageTexts []BACnetOptionalCharacterString
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetOptionalCharacterStringParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'eventMessageTexts' field of BACnetConstructedDataEventMessageTexts")
			}
			eventMessageTexts = append(eventMessageTexts, _item.(BACnetOptionalCharacterString))
		}
	}
	if closeErr := readBuffer.CloseContext("eventMessageTexts", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventMessageTexts")
	}

	// Virtual field
	_toOffnormalText := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTexts[0]) }, func() any { return CastBACnetOptionalCharacterString(nil) }))
	toOffnormalText := _toOffnormalText
	_ = toOffnormalText

	// Virtual field
	_toFaultText := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTexts[1]) }, func() any { return CastBACnetOptionalCharacterString(nil) }))
	toFaultText := _toFaultText
	_ = toFaultText

	// Virtual field
	_toNormalText := CastBACnetOptionalCharacterString(utils.InlineIf(bool((len(eventMessageTexts)) == (3)), func() any { return CastBACnetOptionalCharacterString(eventMessageTexts[2]) }, func() any { return CastBACnetOptionalCharacterString(nil) }))
	toNormalText := _toNormalText
	_ = toNormalText

	// Validation
	if !(bool(bool((arrayIndexArgument) != (nil))) || bool(bool((len(eventMessageTexts)) == (3)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"eventMessageTexts should have exactly 3 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventMessageTexts"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventMessageTexts")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventMessageTexts{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		EventMessageTexts:    eventMessageTexts,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventMessageTexts) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventMessageTexts) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventMessageTexts"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventMessageTexts")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (eventMessageTexts)
		if pushErr := writeBuffer.PushContext("eventMessageTexts", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for eventMessageTexts")
		}
		for _curItem, _element := range m.GetEventMessageTexts() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetEventMessageTexts()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'eventMessageTexts' field")
			}
		}
		if popErr := writeBuffer.PopContext("eventMessageTexts", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for eventMessageTexts")
		}
		// Virtual field
		if _toOffnormalTextErr := writeBuffer.WriteVirtual(ctx, "toOffnormalText", m.GetToOffnormalText()); _toOffnormalTextErr != nil {
			return errors.Wrap(_toOffnormalTextErr, "Error serializing 'toOffnormalText' field")
		}
		// Virtual field
		if _toFaultTextErr := writeBuffer.WriteVirtual(ctx, "toFaultText", m.GetToFaultText()); _toFaultTextErr != nil {
			return errors.Wrap(_toFaultTextErr, "Error serializing 'toFaultText' field")
		}
		// Virtual field
		if _toNormalTextErr := writeBuffer.WriteVirtual(ctx, "toNormalText", m.GetToNormalText()); _toNormalTextErr != nil {
			return errors.Wrap(_toNormalTextErr, "Error serializing 'toNormalText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventMessageTexts"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventMessageTexts")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventMessageTexts) isBACnetConstructedDataEventMessageTexts() bool {
	return true
}

func (m *_BACnetConstructedDataEventMessageTexts) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
