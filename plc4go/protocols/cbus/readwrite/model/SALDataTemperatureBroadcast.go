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

// SALDataTemperatureBroadcast is the corresponding interface of SALDataTemperatureBroadcast
type SALDataTemperatureBroadcast interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetTemperatureBroadcastData returns TemperatureBroadcastData (property field)
	GetTemperatureBroadcastData() TemperatureBroadcastData
}

// SALDataTemperatureBroadcastExactly can be used when we want exactly this type and not a type which fulfills SALDataTemperatureBroadcast.
// This is useful for switch cases.
type SALDataTemperatureBroadcastExactly interface {
	SALDataTemperatureBroadcast
	isSALDataTemperatureBroadcast() bool
}

// _SALDataTemperatureBroadcast is the data-structure of this message
type _SALDataTemperatureBroadcast struct {
	*_SALData
	TemperatureBroadcastData TemperatureBroadcastData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTemperatureBroadcast) GetApplicationId() ApplicationId {
	return ApplicationId_TEMPERATURE_BROADCAST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTemperatureBroadcast) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataTemperatureBroadcast) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataTemperatureBroadcast) GetTemperatureBroadcastData() TemperatureBroadcastData {
	return m.TemperatureBroadcastData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataTemperatureBroadcast factory function for _SALDataTemperatureBroadcast
func NewSALDataTemperatureBroadcast(temperatureBroadcastData TemperatureBroadcastData, salData SALData) *_SALDataTemperatureBroadcast {
	_result := &_SALDataTemperatureBroadcast{
		TemperatureBroadcastData: temperatureBroadcastData,
		_SALData:                 NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTemperatureBroadcast(structType any) SALDataTemperatureBroadcast {
	if casted, ok := structType.(SALDataTemperatureBroadcast); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTemperatureBroadcast); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTemperatureBroadcast) GetTypeName() string {
	return "SALDataTemperatureBroadcast"
}

func (m *_SALDataTemperatureBroadcast) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (temperatureBroadcastData)
	lengthInBits += m.TemperatureBroadcastData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataTemperatureBroadcast) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataTemperatureBroadcastParse(theBytes []byte, applicationId ApplicationId) (SALDataTemperatureBroadcast, error) {
	return SALDataTemperatureBroadcastParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataTemperatureBroadcastParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataTemperatureBroadcast, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTemperatureBroadcast"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTemperatureBroadcast")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (temperatureBroadcastData)
	if pullErr := readBuffer.PullContext("temperatureBroadcastData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for temperatureBroadcastData")
	}
	_temperatureBroadcastData, _temperatureBroadcastDataErr := TemperatureBroadcastDataParseWithBuffer(ctx, readBuffer)
	if _temperatureBroadcastDataErr != nil {
		return nil, errors.Wrap(_temperatureBroadcastDataErr, "Error parsing 'temperatureBroadcastData' field of SALDataTemperatureBroadcast")
	}
	temperatureBroadcastData := _temperatureBroadcastData.(TemperatureBroadcastData)
	if closeErr := readBuffer.CloseContext("temperatureBroadcastData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for temperatureBroadcastData")
	}

	if closeErr := readBuffer.CloseContext("SALDataTemperatureBroadcast"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTemperatureBroadcast")
	}

	// Create a partially initialized instance
	_child := &_SALDataTemperatureBroadcast{
		_SALData:                 &_SALData{},
		TemperatureBroadcastData: temperatureBroadcastData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataTemperatureBroadcast) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTemperatureBroadcast) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTemperatureBroadcast"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTemperatureBroadcast")
		}

		// Simple Field (temperatureBroadcastData)
		if pushErr := writeBuffer.PushContext("temperatureBroadcastData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for temperatureBroadcastData")
		}
		_temperatureBroadcastDataErr := writeBuffer.WriteSerializable(ctx, m.GetTemperatureBroadcastData())
		if popErr := writeBuffer.PopContext("temperatureBroadcastData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for temperatureBroadcastData")
		}
		if _temperatureBroadcastDataErr != nil {
			return errors.Wrap(_temperatureBroadcastDataErr, "Error serializing 'temperatureBroadcastData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataTemperatureBroadcast"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTemperatureBroadcast")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataTemperatureBroadcast) isSALDataTemperatureBroadcast() bool {
	return true
}

func (m *_SALDataTemperatureBroadcast) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
