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

// Constant values.
const S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH uint16 = uint16(28)

// S7PayloadUserDataItemCpuFunctionReadSzlResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionReadSzlResponse
type S7PayloadUserDataItemCpuFunctionReadSzlResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetSzlId returns SzlId (property field)
	GetSzlId() SzlId
	// GetSzlIndex returns SzlIndex (property field)
	GetSzlIndex() uint16
	// GetItems returns Items (property field)
	GetItems() []SzlDataTreeItem
}

// S7PayloadUserDataItemCpuFunctionReadSzlResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionReadSzlResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionReadSzlResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionReadSzlResponse
	isS7PayloadUserDataItemCpuFunctionReadSzlResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionReadSzlResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionReadSzlResponse struct {
	*_S7PayloadUserDataItem
	SzlId    SzlId
	SzlIndex uint16
	Items    []SzlDataTreeItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetCpuSubfunction() uint8 {
	return 0x01
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetSzlId() SzlId {
	return m.SzlId
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetSzlIndex() uint16 {
	return m.SzlIndex
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetItems() []SzlDataTreeItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetSzlItemLength() uint16 {
	return S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionReadSzlResponse factory function for _S7PayloadUserDataItemCpuFunctionReadSzlResponse
func NewS7PayloadUserDataItemCpuFunctionReadSzlResponse(szlId SzlId, szlIndex uint16, items []SzlDataTreeItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		SzlId:                  szlId,
		SzlIndex:               szlIndex,
		Items:                  items,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionReadSzlResponse(structType any) S7PayloadUserDataItemCpuFunctionReadSzlResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionReadSzlResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionReadSzlResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (szlId)
	lengthInBits += m.SzlId.GetLengthInBits(ctx)

	// Simple field (szlIndex)
	lengthInBits += 16

	// Const Field (szlItemLength)
	lengthInBits += 16

	// Implicit Field (szlItemCount)
	lengthInBits += 16

	// Array field
	if len(m.Items) > 0 {
		for _curItem, element := range m.Items {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Items), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParse(theBytes []byte, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionReadSzlResponse, error) {
	return S7PayloadUserDataItemCpuFunctionReadSzlResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionReadSzlResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionReadSzlResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (szlId)
	if pullErr := readBuffer.PullContext("szlId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for szlId")
	}
	_szlId, _szlIdErr := SzlIdParseWithBuffer(ctx, readBuffer)
	if _szlIdErr != nil {
		return nil, errors.Wrap(_szlIdErr, "Error parsing 'szlId' field of S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}
	szlId := _szlId.(SzlId)
	if closeErr := readBuffer.CloseContext("szlId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for szlId")
	}

	// Simple Field (szlIndex)
	_szlIndex, _szlIndexErr := readBuffer.ReadUint16("szlIndex", 16)
	if _szlIndexErr != nil {
		return nil, errors.Wrap(_szlIndexErr, "Error parsing 'szlIndex' field of S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}
	szlIndex := _szlIndex

	// Const Field (szlItemLength)
	szlItemLength, _szlItemLengthErr := readBuffer.ReadUint16("szlItemLength", 16)
	if _szlItemLengthErr != nil {
		return nil, errors.Wrap(_szlItemLengthErr, "Error parsing 'szlItemLength' field of S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}
	if szlItemLength != S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionReadSzlResponse_SZLITEMLENGTH) + " but got " + fmt.Sprintf("%d", szlItemLength))
	}

	// Implicit Field (szlItemCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	szlItemCount, _szlItemCountErr := readBuffer.ReadUint16("szlItemCount", 16)
	_ = szlItemCount
	if _szlItemCountErr != nil {
		return nil, errors.Wrap(_szlItemCountErr, "Error parsing 'szlItemCount' field of S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for items")
	}
	// Count array
	items := make([]SzlDataTreeItem, szlItemCount)
	// This happens when the size is set conditional to 0
	if len(items) == 0 {
		items = nil
	}
	{
		_numItems := uint16(szlItemCount)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := SzlDataTreeItemParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field of S7PayloadUserDataItemCpuFunctionReadSzlResponse")
			}
			items[_curItem] = _item.(SzlDataTreeItem)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for items")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionReadSzlResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		SzlId:                  szlId,
		SzlIndex:               szlIndex,
		Items:                  items,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
		}

		// Simple Field (szlId)
		if pushErr := writeBuffer.PushContext("szlId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for szlId")
		}
		_szlIdErr := writeBuffer.WriteSerializable(ctx, m.GetSzlId())
		if popErr := writeBuffer.PopContext("szlId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for szlId")
		}
		if _szlIdErr != nil {
			return errors.Wrap(_szlIdErr, "Error serializing 'szlId' field")
		}

		// Simple Field (szlIndex)
		szlIndex := uint16(m.GetSzlIndex())
		_szlIndexErr := writeBuffer.WriteUint16("szlIndex", 16, (szlIndex))
		if _szlIndexErr != nil {
			return errors.Wrap(_szlIndexErr, "Error serializing 'szlIndex' field")
		}

		// Const Field (szlItemLength)
		_szlItemLengthErr := writeBuffer.WriteUint16("szlItemLength", 16, 28)
		if _szlItemLengthErr != nil {
			return errors.Wrap(_szlItemLengthErr, "Error serializing 'szlItemLength' field")
		}

		// Implicit Field (szlItemCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		szlItemCount := uint16(uint16(len(m.GetItems())))
		_szlItemCountErr := writeBuffer.WriteUint16("szlItemCount", 16, (szlItemCount))
		if _szlItemCountErr != nil {
			return errors.Wrap(_szlItemCountErr, "Error serializing 'szlItemCount' field")
		}

		// Array Field (items)
		if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for items")
		}
		for _curItem, _element := range m.GetItems() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItems()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'items' field")
			}
		}
		if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for items")
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionReadSzlResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionReadSzlResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) isS7PayloadUserDataItemCpuFunctionReadSzlResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionReadSzlResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
