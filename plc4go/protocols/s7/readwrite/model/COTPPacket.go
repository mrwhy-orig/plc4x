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

// COTPPacket is the corresponding interface of COTPPacket
type COTPPacket interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTpduCode returns TpduCode (discriminator field)
	GetTpduCode() uint8
	// GetParameters returns Parameters (property field)
	GetParameters() []COTPParameter
	// GetPayload returns Payload (property field)
	GetPayload() S7Message
}

// COTPPacketExactly can be used when we want exactly this type and not a type which fulfills COTPPacket.
// This is useful for switch cases.
type COTPPacketExactly interface {
	COTPPacket
	isCOTPPacket() bool
}

// _COTPPacket is the data-structure of this message
type _COTPPacket struct {
	_COTPPacketChildRequirements
	Parameters []COTPParameter
	Payload    S7Message

	// Arguments.
	CotpLen uint16
}

type _COTPPacketChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetTpduCode() uint8
}

type COTPPacketParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type COTPPacketChild interface {
	utils.Serializable
	InitializeParent(parent COTPPacket, parameters []COTPParameter, payload S7Message)
	GetParent() *COTPPacket

	GetTypeName() string
	COTPPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacket) GetParameters() []COTPParameter {
	return m.Parameters
}

func (m *_COTPPacket) GetPayload() S7Message {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCOTPPacket factory function for _COTPPacket
func NewCOTPPacket(parameters []COTPParameter, payload S7Message, cotpLen uint16) *_COTPPacket {
	return &_COTPPacket{Parameters: parameters, Payload: payload, CotpLen: cotpLen}
}

// Deprecated: use the interface for direct cast
func CastCOTPPacket(structType any) COTPPacket {
	if casted, ok := structType.(COTPPacket); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacket); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacket) GetTypeName() string {
	return "COTPPacket"
}

func (m *_COTPPacket) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8
	// Discriminator Field (tpduCode)
	lengthInBits += 8

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += m.Payload.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_COTPPacket) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPPacketParse(theBytes []byte, cotpLen uint16) (COTPPacket, error) {
	return COTPPacketParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cotpLen)
}

func COTPPacketParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacket, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacket"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacket")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	var startPos = positionAware.GetPos()
	_ = startPos

	// Implicit Field (headerLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	headerLength, _headerLengthErr := readBuffer.ReadUint8("headerLength", 8)
	_ = headerLength
	if _headerLengthErr != nil {
		return nil, errors.Wrap(_headerLengthErr, "Error parsing 'headerLength' field of COTPPacket")
	}

	// Discriminator Field (tpduCode) (Used as input to a switch field)
	tpduCode, _tpduCodeErr := readBuffer.ReadUint8("tpduCode", 8)
	if _tpduCodeErr != nil {
		return nil, errors.Wrap(_tpduCodeErr, "Error parsing 'tpduCode' field of COTPPacket")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type COTPPacketChildSerializeRequirement interface {
		COTPPacket
		InitializeParent(COTPPacket, []COTPParameter, S7Message)
		GetParent() COTPPacket
	}
	var _childTemp any
	var _child COTPPacketChildSerializeRequirement
	var typeSwitchError error
	switch {
	case tpduCode == 0xF0: // COTPPacketData
		_childTemp, typeSwitchError = COTPPacketDataParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0xE0: // COTPPacketConnectionRequest
		_childTemp, typeSwitchError = COTPPacketConnectionRequestParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0xD0: // COTPPacketConnectionResponse
		_childTemp, typeSwitchError = COTPPacketConnectionResponseParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0x80: // COTPPacketDisconnectRequest
		_childTemp, typeSwitchError = COTPPacketDisconnectRequestParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0xC0: // COTPPacketDisconnectResponse
		_childTemp, typeSwitchError = COTPPacketDisconnectResponseParseWithBuffer(ctx, readBuffer, cotpLen)
	case tpduCode == 0x70: // COTPPacketTpduError
		_childTemp, typeSwitchError = COTPPacketTpduErrorParseWithBuffer(ctx, readBuffer, cotpLen)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [tpduCode=%v]", tpduCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of COTPPacket")
	}
	_child = _childTemp.(COTPPacketChildSerializeRequirement)

	// Array field (parameters)
	if pullErr := readBuffer.PullContext("parameters", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for parameters")
	}
	// Length array
	var parameters []COTPParameter
	{
		_parametersLength := uint16((uint16(headerLength) + uint16(uint16(1)))) - uint16((positionAware.GetPos() - startPos))
		_parametersEndPos := positionAware.GetPos() + uint16(_parametersLength)
		for positionAware.GetPos() < _parametersEndPos {
			_item, _err := COTPParameterParseWithBuffer(ctx, readBuffer, uint8((uint8(headerLength)+uint8(uint8(1))))-uint8((positionAware.GetPos()-startPos)))
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'parameters' field of COTPPacket")
			}
			parameters = append(parameters, _item.(COTPParameter))
		}
	}
	if closeErr := readBuffer.CloseContext("parameters", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for parameters")
	}

	// Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
	var payload S7Message = nil
	if bool((positionAware.GetPos() - startPos) < (cotpLen)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for payload")
		}
		_val, _err := S7MessageParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'payload' field of COTPPacket")
		default:
			payload = _val.(S7Message)
			if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for payload")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("COTPPacket"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacket")
	}

	// Finish initializing
	_child.InitializeParent(_child, parameters, payload)
	return _child, nil
}

func (pm *_COTPPacket) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPPacket, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("COTPPacket"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for COTPPacket")
	}

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8((uint8((utils.InlineIf((bool((m.GetPayload()) != (nil))), func() any { return uint8((m.GetPayload()).GetLengthInBytes(ctx)) }, func() any { return uint8(uint8(0)) }).(uint8))) + uint8(uint8(1)))))
	_headerLengthErr := writeBuffer.WriteUint8("headerLength", 8, (headerLength))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	// Discriminator Field (tpduCode) (Used as input to a switch field)
	tpduCode := uint8(child.GetTpduCode())
	_tpduCodeErr := writeBuffer.WriteUint8("tpduCode", 8, (tpduCode))

	if _tpduCodeErr != nil {
		return errors.Wrap(_tpduCodeErr, "Error serializing 'tpduCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Array Field (parameters)
	if pushErr := writeBuffer.PushContext("parameters", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for parameters")
	}
	for _curItem, _element := range m.GetParameters() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetParameters()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'parameters' field")
		}
	}
	if popErr := writeBuffer.PopContext("parameters", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for parameters")
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload S7Message = nil
	if m.GetPayload() != nil {
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		payload = m.GetPayload()
		_payloadErr := writeBuffer.WriteSerializable(ctx, payload)
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	if popErr := writeBuffer.PopContext("COTPPacket"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for COTPPacket")
	}
	return nil
}

////
// Arguments Getter

func (m *_COTPPacket) GetCotpLen() uint16 {
	return m.CotpLen
}

//
////

func (m *_COTPPacket) isCOTPPacket() bool {
	return true
}

func (m *_COTPPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
