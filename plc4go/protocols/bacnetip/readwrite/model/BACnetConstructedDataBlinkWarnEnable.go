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

// BACnetConstructedDataBlinkWarnEnable is the corresponding interface of BACnetConstructedDataBlinkWarnEnable
type BACnetConstructedDataBlinkWarnEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBlinkWarnEnable returns BlinkWarnEnable (property field)
	GetBlinkWarnEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataBlinkWarnEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBlinkWarnEnable.
// This is useful for switch cases.
type BACnetConstructedDataBlinkWarnEnableExactly interface {
	BACnetConstructedDataBlinkWarnEnable
	isBACnetConstructedDataBlinkWarnEnable() bool
}

// _BACnetConstructedDataBlinkWarnEnable is the data-structure of this message
type _BACnetConstructedDataBlinkWarnEnable struct {
	*_BACnetConstructedData
	BlinkWarnEnable BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BLINK_WARN_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetBlinkWarnEnable() BACnetApplicationTagBoolean {
	return m.BlinkWarnEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBlinkWarnEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetBlinkWarnEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBlinkWarnEnable factory function for _BACnetConstructedDataBlinkWarnEnable
func NewBACnetConstructedDataBlinkWarnEnable(blinkWarnEnable BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBlinkWarnEnable {
	_result := &_BACnetConstructedDataBlinkWarnEnable{
		BlinkWarnEnable:        blinkWarnEnable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBlinkWarnEnable(structType any) BACnetConstructedDataBlinkWarnEnable {
	if casted, ok := structType.(BACnetConstructedDataBlinkWarnEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBlinkWarnEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetTypeName() string {
	return "BACnetConstructedDataBlinkWarnEnable"
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (blinkWarnEnable)
	lengthInBits += m.BlinkWarnEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBlinkWarnEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBlinkWarnEnableParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBlinkWarnEnable, error) {
	return BACnetConstructedDataBlinkWarnEnableParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBlinkWarnEnableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBlinkWarnEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBlinkWarnEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBlinkWarnEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (blinkWarnEnable)
	if pullErr := readBuffer.PullContext("blinkWarnEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for blinkWarnEnable")
	}
	_blinkWarnEnable, _blinkWarnEnableErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _blinkWarnEnableErr != nil {
		return nil, errors.Wrap(_blinkWarnEnableErr, "Error parsing 'blinkWarnEnable' field of BACnetConstructedDataBlinkWarnEnable")
	}
	blinkWarnEnable := _blinkWarnEnable.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("blinkWarnEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for blinkWarnEnable")
	}

	// Virtual field
	_actualValue := blinkWarnEnable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBlinkWarnEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBlinkWarnEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBlinkWarnEnable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BlinkWarnEnable: blinkWarnEnable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBlinkWarnEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBlinkWarnEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBlinkWarnEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBlinkWarnEnable")
		}

		// Simple Field (blinkWarnEnable)
		if pushErr := writeBuffer.PushContext("blinkWarnEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for blinkWarnEnable")
		}
		_blinkWarnEnableErr := writeBuffer.WriteSerializable(ctx, m.GetBlinkWarnEnable())
		if popErr := writeBuffer.PopContext("blinkWarnEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for blinkWarnEnable")
		}
		if _blinkWarnEnableErr != nil {
			return errors.Wrap(_blinkWarnEnableErr, "Error serializing 'blinkWarnEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBlinkWarnEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBlinkWarnEnable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBlinkWarnEnable) isBACnetConstructedDataBlinkWarnEnable() bool {
	return true
}

func (m *_BACnetConstructedDataBlinkWarnEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
