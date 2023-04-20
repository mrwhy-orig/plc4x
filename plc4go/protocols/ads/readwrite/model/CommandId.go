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

// CommandId is an enum
type CommandId uint16

type ICommandId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	CommandId_INVALID                        CommandId = 0x0000
	CommandId_ADS_READ_DEVICE_INFO           CommandId = 0x0001
	CommandId_ADS_READ                       CommandId = 0x0002
	CommandId_ADS_WRITE                      CommandId = 0x0003
	CommandId_ADS_READ_STATE                 CommandId = 0x0004
	CommandId_ADS_WRITE_CONTROL              CommandId = 0x0005
	CommandId_ADS_ADD_DEVICE_NOTIFICATION    CommandId = 0x0006
	CommandId_ADS_DELETE_DEVICE_NOTIFICATION CommandId = 0x0007
	CommandId_ADS_DEVICE_NOTIFICATION        CommandId = 0x0008
	CommandId_ADS_READ_WRITE                 CommandId = 0x0009
)

var CommandIdValues []CommandId

func init() {
	_ = errors.New
	CommandIdValues = []CommandId{
		CommandId_INVALID,
		CommandId_ADS_READ_DEVICE_INFO,
		CommandId_ADS_READ,
		CommandId_ADS_WRITE,
		CommandId_ADS_READ_STATE,
		CommandId_ADS_WRITE_CONTROL,
		CommandId_ADS_ADD_DEVICE_NOTIFICATION,
		CommandId_ADS_DELETE_DEVICE_NOTIFICATION,
		CommandId_ADS_DEVICE_NOTIFICATION,
		CommandId_ADS_READ_WRITE,
	}
}

func CommandIdByValue(value uint16) (enum CommandId, ok bool) {
	switch value {
	case 0x0000:
		return CommandId_INVALID, true
	case 0x0001:
		return CommandId_ADS_READ_DEVICE_INFO, true
	case 0x0002:
		return CommandId_ADS_READ, true
	case 0x0003:
		return CommandId_ADS_WRITE, true
	case 0x0004:
		return CommandId_ADS_READ_STATE, true
	case 0x0005:
		return CommandId_ADS_WRITE_CONTROL, true
	case 0x0006:
		return CommandId_ADS_ADD_DEVICE_NOTIFICATION, true
	case 0x0007:
		return CommandId_ADS_DELETE_DEVICE_NOTIFICATION, true
	case 0x0008:
		return CommandId_ADS_DEVICE_NOTIFICATION, true
	case 0x0009:
		return CommandId_ADS_READ_WRITE, true
	}
	return 0, false
}

func CommandIdByName(value string) (enum CommandId, ok bool) {
	switch value {
	case "INVALID":
		return CommandId_INVALID, true
	case "ADS_READ_DEVICE_INFO":
		return CommandId_ADS_READ_DEVICE_INFO, true
	case "ADS_READ":
		return CommandId_ADS_READ, true
	case "ADS_WRITE":
		return CommandId_ADS_WRITE, true
	case "ADS_READ_STATE":
		return CommandId_ADS_READ_STATE, true
	case "ADS_WRITE_CONTROL":
		return CommandId_ADS_WRITE_CONTROL, true
	case "ADS_ADD_DEVICE_NOTIFICATION":
		return CommandId_ADS_ADD_DEVICE_NOTIFICATION, true
	case "ADS_DELETE_DEVICE_NOTIFICATION":
		return CommandId_ADS_DELETE_DEVICE_NOTIFICATION, true
	case "ADS_DEVICE_NOTIFICATION":
		return CommandId_ADS_DEVICE_NOTIFICATION, true
	case "ADS_READ_WRITE":
		return CommandId_ADS_READ_WRITE, true
	}
	return 0, false
}

func CommandIdKnows(value uint16) bool {
	for _, typeValue := range CommandIdValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCommandId(structType any) CommandId {
	castFunc := func(typ any) CommandId {
		if sCommandId, ok := typ.(CommandId); ok {
			return sCommandId
		}
		return 0
	}
	return castFunc(structType)
}

func (m CommandId) GetLengthInBits(ctx context.Context) uint16 {
	return 16
}

func (m CommandId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CommandIdParse(ctx context.Context, theBytes []byte) (CommandId, error) {
	return CommandIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CommandIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CommandId, error) {
	val, err := readBuffer.ReadUint16("CommandId", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CommandId")
	}
	if enum, ok := CommandIdByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return CommandId(val), nil
	} else {
		return enum, nil
	}
}

func (e CommandId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CommandId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("CommandId", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CommandId) PLC4XEnumName() string {
	switch e {
	case CommandId_INVALID:
		return "INVALID"
	case CommandId_ADS_READ_DEVICE_INFO:
		return "ADS_READ_DEVICE_INFO"
	case CommandId_ADS_READ:
		return "ADS_READ"
	case CommandId_ADS_WRITE:
		return "ADS_WRITE"
	case CommandId_ADS_READ_STATE:
		return "ADS_READ_STATE"
	case CommandId_ADS_WRITE_CONTROL:
		return "ADS_WRITE_CONTROL"
	case CommandId_ADS_ADD_DEVICE_NOTIFICATION:
		return "ADS_ADD_DEVICE_NOTIFICATION"
	case CommandId_ADS_DELETE_DEVICE_NOTIFICATION:
		return "ADS_DELETE_DEVICE_NOTIFICATION"
	case CommandId_ADS_DEVICE_NOTIFICATION:
		return "ADS_DEVICE_NOTIFICATION"
	case CommandId_ADS_READ_WRITE:
		return "ADS_READ_WRITE"
	}
	return ""
}

func (e CommandId) String() string {
	return e.PLC4XEnumName()
}
