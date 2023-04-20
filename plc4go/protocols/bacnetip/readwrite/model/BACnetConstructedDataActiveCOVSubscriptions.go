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

// BACnetConstructedDataActiveCOVSubscriptions is the corresponding interface of BACnetConstructedDataActiveCOVSubscriptions
type BACnetConstructedDataActiveCOVSubscriptions interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActiveCOVSubscriptions returns ActiveCOVSubscriptions (property field)
	GetActiveCOVSubscriptions() []BACnetCOVSubscription
}

// BACnetConstructedDataActiveCOVSubscriptionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActiveCOVSubscriptions.
// This is useful for switch cases.
type BACnetConstructedDataActiveCOVSubscriptionsExactly interface {
	BACnetConstructedDataActiveCOVSubscriptions
	isBACnetConstructedDataActiveCOVSubscriptions() bool
}

// _BACnetConstructedDataActiveCOVSubscriptions is the data-structure of this message
type _BACnetConstructedDataActiveCOVSubscriptions struct {
	*_BACnetConstructedData
	ActiveCOVSubscriptions []BACnetCOVSubscription
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_COV_SUBSCRIPTIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActiveCOVSubscriptions) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetActiveCOVSubscriptions() []BACnetCOVSubscription {
	return m.ActiveCOVSubscriptions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveCOVSubscriptions factory function for _BACnetConstructedDataActiveCOVSubscriptions
func NewBACnetConstructedDataActiveCOVSubscriptions(activeCOVSubscriptions []BACnetCOVSubscription, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActiveCOVSubscriptions {
	_result := &_BACnetConstructedDataActiveCOVSubscriptions{
		ActiveCOVSubscriptions: activeCOVSubscriptions,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActiveCOVSubscriptions(structType any) BACnetConstructedDataActiveCOVSubscriptions {
	if casted, ok := structType.(BACnetConstructedDataActiveCOVSubscriptions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveCOVSubscriptions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetTypeName() string {
	return "BACnetConstructedDataActiveCOVSubscriptions"
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.ActiveCOVSubscriptions) > 0 {
		for _, element := range m.ActiveCOVSubscriptions {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataActiveCOVSubscriptionsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActiveCOVSubscriptions, error) {
	return BACnetConstructedDataActiveCOVSubscriptionsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataActiveCOVSubscriptionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActiveCOVSubscriptions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveCOVSubscriptions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveCOVSubscriptions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (activeCOVSubscriptions)
	if pullErr := readBuffer.PullContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for activeCOVSubscriptions")
	}
	// Terminated array
	var activeCOVSubscriptions []BACnetCOVSubscription
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCOVSubscriptionParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'activeCOVSubscriptions' field of BACnetConstructedDataActiveCOVSubscriptions")
			}
			activeCOVSubscriptions = append(activeCOVSubscriptions, _item.(BACnetCOVSubscription))
		}
	}
	if closeErr := readBuffer.CloseContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for activeCOVSubscriptions")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveCOVSubscriptions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveCOVSubscriptions")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataActiveCOVSubscriptions{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ActiveCOVSubscriptions: activeCOVSubscriptions,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveCOVSubscriptions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveCOVSubscriptions")
		}

		// Array Field (activeCOVSubscriptions)
		if pushErr := writeBuffer.PushContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for activeCOVSubscriptions")
		}
		for _curItem, _element := range m.GetActiveCOVSubscriptions() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetActiveCOVSubscriptions()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'activeCOVSubscriptions' field")
			}
		}
		if popErr := writeBuffer.PopContext("activeCOVSubscriptions", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for activeCOVSubscriptions")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveCOVSubscriptions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveCOVSubscriptions")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) isBACnetConstructedDataActiveCOVSubscriptions() bool {
	return true
}

func (m *_BACnetConstructedDataActiveCOVSubscriptions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
