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

// ModbusPDUReadDiscreteInputsRequest is the corresponding interface of ModbusPDUReadDiscreteInputsRequest
type ModbusPDUReadDiscreteInputsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
}

// ModbusPDUReadDiscreteInputsRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadDiscreteInputsRequest.
// This is useful for switch cases.
type ModbusPDUReadDiscreteInputsRequestExactly interface {
	ModbusPDUReadDiscreteInputsRequest
	isModbusPDUReadDiscreteInputsRequest() bool
}

// _ModbusPDUReadDiscreteInputsRequest is the data-structure of this message
type _ModbusPDUReadDiscreteInputsRequest struct {
	*_ModbusPDU
	StartingAddress uint16
	Quantity        uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadDiscreteInputsRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetFunctionFlag() uint8 {
	return 0x02
}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadDiscreteInputsRequest) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadDiscreteInputsRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetQuantity() uint16 {
	return m.Quantity
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDiscreteInputsRequest factory function for _ModbusPDUReadDiscreteInputsRequest
func NewModbusPDUReadDiscreteInputsRequest(startingAddress uint16, quantity uint16) *_ModbusPDUReadDiscreteInputsRequest {
	_result := &_ModbusPDUReadDiscreteInputsRequest{
		StartingAddress: startingAddress,
		Quantity:        quantity,
		_ModbusPDU:      NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadDiscreteInputsRequest(structType any) ModbusPDUReadDiscreteInputsRequest {
	if casted, ok := structType.(ModbusPDUReadDiscreteInputsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadDiscreteInputsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetTypeName() string {
	return "ModbusPDUReadDiscreteInputsRequest"
}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (startingAddress)
	lengthInBits += 16

	// Simple field (quantity)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUReadDiscreteInputsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReadDiscreteInputsRequestParse(theBytes []byte, response bool) (ModbusPDUReadDiscreteInputsRequest, error) {
	return ModbusPDUReadDiscreteInputsRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadDiscreteInputsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadDiscreteInputsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadDiscreteInputsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadDiscreteInputsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
	_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field of ModbusPDUReadDiscreteInputsRequest")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
	_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field of ModbusPDUReadDiscreteInputsRequest")
	}
	quantity := _quantity

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDiscreteInputsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadDiscreteInputsRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadDiscreteInputsRequest{
		_ModbusPDU:      &_ModbusPDU{},
		StartingAddress: startingAddress,
		Quantity:        quantity,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadDiscreteInputsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadDiscreteInputsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDiscreteInputsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadDiscreteInputsRequest")
		}

		// Simple Field (startingAddress)
		startingAddress := uint16(m.GetStartingAddress())
		_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
		if _startingAddressErr != nil {
			return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
		}

		// Simple Field (quantity)
		quantity := uint16(m.GetQuantity())
		_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
		if _quantityErr != nil {
			return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDiscreteInputsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadDiscreteInputsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUReadDiscreteInputsRequest) isModbusPDUReadDiscreteInputsRequest() bool {
	return true
}

func (m *_ModbusPDUReadDiscreteInputsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
