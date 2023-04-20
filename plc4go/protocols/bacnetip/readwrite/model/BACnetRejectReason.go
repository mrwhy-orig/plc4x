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

// BACnetRejectReason is an enum
type BACnetRejectReason uint8

type IBACnetRejectReason interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetRejectReason_OTHER                       BACnetRejectReason = 0x0
	BACnetRejectReason_BUFFER_OVERFLOW             BACnetRejectReason = 0x1
	BACnetRejectReason_INCONSISTENT_PARAMETERS     BACnetRejectReason = 0x2
	BACnetRejectReason_INVALID_PARAMETER_DATA_TYPE BACnetRejectReason = 0x3
	BACnetRejectReason_INVALID_TAG                 BACnetRejectReason = 0x4
	BACnetRejectReason_MISSING_REQUIRED_PARAMETER  BACnetRejectReason = 0x5
	BACnetRejectReason_PARAMETER_OUT_OF_RANGE      BACnetRejectReason = 0x6
	BACnetRejectReason_TOO_MANY_ARGUMENTS          BACnetRejectReason = 0x7
	BACnetRejectReason_UNDEFINED_ENUMERATION       BACnetRejectReason = 0x8
	BACnetRejectReason_UNRECOGNIZED_SERVICE        BACnetRejectReason = 0x9
	BACnetRejectReason_VENDOR_PROPRIETARY_VALUE    BACnetRejectReason = 0xFF
)

var BACnetRejectReasonValues []BACnetRejectReason

func init() {
	_ = errors.New
	BACnetRejectReasonValues = []BACnetRejectReason{
		BACnetRejectReason_OTHER,
		BACnetRejectReason_BUFFER_OVERFLOW,
		BACnetRejectReason_INCONSISTENT_PARAMETERS,
		BACnetRejectReason_INVALID_PARAMETER_DATA_TYPE,
		BACnetRejectReason_INVALID_TAG,
		BACnetRejectReason_MISSING_REQUIRED_PARAMETER,
		BACnetRejectReason_PARAMETER_OUT_OF_RANGE,
		BACnetRejectReason_TOO_MANY_ARGUMENTS,
		BACnetRejectReason_UNDEFINED_ENUMERATION,
		BACnetRejectReason_UNRECOGNIZED_SERVICE,
		BACnetRejectReason_VENDOR_PROPRIETARY_VALUE,
	}
}

func BACnetRejectReasonByValue(value uint8) (enum BACnetRejectReason, ok bool) {
	switch value {
	case 0x0:
		return BACnetRejectReason_OTHER, true
	case 0x1:
		return BACnetRejectReason_BUFFER_OVERFLOW, true
	case 0x2:
		return BACnetRejectReason_INCONSISTENT_PARAMETERS, true
	case 0x3:
		return BACnetRejectReason_INVALID_PARAMETER_DATA_TYPE, true
	case 0x4:
		return BACnetRejectReason_INVALID_TAG, true
	case 0x5:
		return BACnetRejectReason_MISSING_REQUIRED_PARAMETER, true
	case 0x6:
		return BACnetRejectReason_PARAMETER_OUT_OF_RANGE, true
	case 0x7:
		return BACnetRejectReason_TOO_MANY_ARGUMENTS, true
	case 0x8:
		return BACnetRejectReason_UNDEFINED_ENUMERATION, true
	case 0x9:
		return BACnetRejectReason_UNRECOGNIZED_SERVICE, true
	case 0xFF:
		return BACnetRejectReason_VENDOR_PROPRIETARY_VALUE, true
	}
	return 0, false
}

func BACnetRejectReasonByName(value string) (enum BACnetRejectReason, ok bool) {
	switch value {
	case "OTHER":
		return BACnetRejectReason_OTHER, true
	case "BUFFER_OVERFLOW":
		return BACnetRejectReason_BUFFER_OVERFLOW, true
	case "INCONSISTENT_PARAMETERS":
		return BACnetRejectReason_INCONSISTENT_PARAMETERS, true
	case "INVALID_PARAMETER_DATA_TYPE":
		return BACnetRejectReason_INVALID_PARAMETER_DATA_TYPE, true
	case "INVALID_TAG":
		return BACnetRejectReason_INVALID_TAG, true
	case "MISSING_REQUIRED_PARAMETER":
		return BACnetRejectReason_MISSING_REQUIRED_PARAMETER, true
	case "PARAMETER_OUT_OF_RANGE":
		return BACnetRejectReason_PARAMETER_OUT_OF_RANGE, true
	case "TOO_MANY_ARGUMENTS":
		return BACnetRejectReason_TOO_MANY_ARGUMENTS, true
	case "UNDEFINED_ENUMERATION":
		return BACnetRejectReason_UNDEFINED_ENUMERATION, true
	case "UNRECOGNIZED_SERVICE":
		return BACnetRejectReason_UNRECOGNIZED_SERVICE, true
	case "VENDOR_PROPRIETARY_VALUE":
		return BACnetRejectReason_VENDOR_PROPRIETARY_VALUE, true
	}
	return 0, false
}

func BACnetRejectReasonKnows(value uint8) bool {
	for _, typeValue := range BACnetRejectReasonValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetRejectReason(structType any) BACnetRejectReason {
	castFunc := func(typ any) BACnetRejectReason {
		if sBACnetRejectReason, ok := typ.(BACnetRejectReason); ok {
			return sBACnetRejectReason
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetRejectReason) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetRejectReason) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRejectReasonParse(ctx context.Context, theBytes []byte) (BACnetRejectReason, error) {
	return BACnetRejectReasonParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRejectReasonParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRejectReason, error) {
	val, err := readBuffer.ReadUint8("BACnetRejectReason", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetRejectReason")
	}
	if enum, ok := BACnetRejectReasonByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetRejectReason(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetRejectReason) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetRejectReason) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetRejectReason", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetRejectReason) PLC4XEnumName() string {
	switch e {
	case BACnetRejectReason_OTHER:
		return "OTHER"
	case BACnetRejectReason_BUFFER_OVERFLOW:
		return "BUFFER_OVERFLOW"
	case BACnetRejectReason_INCONSISTENT_PARAMETERS:
		return "INCONSISTENT_PARAMETERS"
	case BACnetRejectReason_INVALID_PARAMETER_DATA_TYPE:
		return "INVALID_PARAMETER_DATA_TYPE"
	case BACnetRejectReason_INVALID_TAG:
		return "INVALID_TAG"
	case BACnetRejectReason_MISSING_REQUIRED_PARAMETER:
		return "MISSING_REQUIRED_PARAMETER"
	case BACnetRejectReason_PARAMETER_OUT_OF_RANGE:
		return "PARAMETER_OUT_OF_RANGE"
	case BACnetRejectReason_TOO_MANY_ARGUMENTS:
		return "TOO_MANY_ARGUMENTS"
	case BACnetRejectReason_UNDEFINED_ENUMERATION:
		return "UNDEFINED_ENUMERATION"
	case BACnetRejectReason_UNRECOGNIZED_SERVICE:
		return "UNRECOGNIZED_SERVICE"
	case BACnetRejectReason_VENDOR_PROPRIETARY_VALUE:
		return "VENDOR_PROPRIETARY_VALUE"
	}
	return ""
}

func (e BACnetRejectReason) String() string {
	return e.PLC4XEnumName()
}
