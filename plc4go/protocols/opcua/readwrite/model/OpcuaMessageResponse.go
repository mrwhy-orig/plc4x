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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaMessageResponse is the corresponding interface of OpcuaMessageResponse
type OpcuaMessageResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MessagePDU
	// GetSecurityHeader returns SecurityHeader (property field)
	GetSecurityHeader() SecurityHeader
	// GetMessage returns Message (property field)
	GetMessage() Payload
}

// OpcuaMessageResponseExactly can be used when we want exactly this type and not a type which fulfills OpcuaMessageResponse.
// This is useful for switch cases.
type OpcuaMessageResponseExactly interface {
	OpcuaMessageResponse
	isOpcuaMessageResponse() bool
}

// _OpcuaMessageResponse is the data-structure of this message
type _OpcuaMessageResponse struct {
	*_MessagePDU
	SecurityHeader SecurityHeader
	Message        Payload

	// Arguments.
	TotalLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaMessageResponse) GetMessageType() string {
	return "MSG"
}

func (m *_OpcuaMessageResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaMessageResponse) InitializeParent(parent MessagePDU, chunk ChunkType) {
	m.Chunk = chunk
}

func (m *_OpcuaMessageResponse) GetParent() MessagePDU {
	return m._MessagePDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaMessageResponse) GetSecurityHeader() SecurityHeader {
	return m.SecurityHeader
}

func (m *_OpcuaMessageResponse) GetMessage() Payload {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewOpcuaMessageResponse factory function for _OpcuaMessageResponse
func NewOpcuaMessageResponse(securityHeader SecurityHeader, message Payload, chunk ChunkType, totalLength uint32) *_OpcuaMessageResponse {
	_result := &_OpcuaMessageResponse{
		SecurityHeader: securityHeader,
		Message:        message,
		_MessagePDU:    NewMessagePDU(chunk),
	}
	_result._MessagePDU._MessagePDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastOpcuaMessageResponse(structType any) OpcuaMessageResponse {
	if casted, ok := structType.(OpcuaMessageResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaMessageResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaMessageResponse) GetTypeName() string {
	return "OpcuaMessageResponse"
}

func (m *_OpcuaMessageResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (securityHeader)
	lengthInBits += m.SecurityHeader.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaMessageResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaMessageResponseParse(ctx context.Context, theBytes []byte, totalLength uint32, response bool) (OpcuaMessageResponse, error) {
	return OpcuaMessageResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), totalLength, response)
}

func OpcuaMessageResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, totalLength uint32, response bool) (OpcuaMessageResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("OpcuaMessageResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaMessageResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securityHeader)
	if pullErr := readBuffer.PullContext("securityHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityHeader")
	}
	_securityHeader, _securityHeaderErr := SecurityHeaderParseWithBuffer(ctx, readBuffer)
	if _securityHeaderErr != nil {
		return nil, errors.Wrap(_securityHeaderErr, "Error parsing 'securityHeader' field of OpcuaMessageResponse")
	}
	securityHeader := _securityHeader.(SecurityHeader)
	if closeErr := readBuffer.CloseContext("securityHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityHeader")
	}

	// Simple Field (message)
	if pullErr := readBuffer.PullContext("message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for message")
	}
	_message, _messageErr := PayloadParseWithBuffer(ctx, readBuffer, bool(bool(false)), uint32(uint32(uint32(totalLength)-uint32(securityHeader.GetLengthInBytes(ctx)))-uint32(uint32(16))))
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field of OpcuaMessageResponse")
	}
	message := _message.(Payload)
	if closeErr := readBuffer.CloseContext("message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for message")
	}

	if closeErr := readBuffer.CloseContext("OpcuaMessageResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaMessageResponse")
	}

	// Create a partially initialized instance
	_child := &_OpcuaMessageResponse{
		_MessagePDU:    &_MessagePDU{},
		SecurityHeader: securityHeader,
		Message:        message,
	}
	_child._MessagePDU._MessagePDUChildRequirements = _child
	return _child, nil
}

func (m *_OpcuaMessageResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaMessageResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaMessageResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaMessageResponse")
		}

		// Simple Field (securityHeader)
		if pushErr := writeBuffer.PushContext("securityHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityHeader")
		}
		_securityHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityHeader())
		if popErr := writeBuffer.PopContext("securityHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityHeader")
		}
		if _securityHeaderErr != nil {
			return errors.Wrap(_securityHeaderErr, "Error serializing 'securityHeader' field")
		}

		// Simple Field (message)
		if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for message")
		}
		_messageErr := writeBuffer.WriteSerializable(ctx, m.GetMessage())
		if popErr := writeBuffer.PopContext("message"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for message")
		}
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaMessageResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaMessageResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_OpcuaMessageResponse) GetTotalLength() uint32 {
	return m.TotalLength
}

//
////

func (m *_OpcuaMessageResponse) isOpcuaMessageResponse() bool {
	return true
}

func (m *_OpcuaMessageResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
