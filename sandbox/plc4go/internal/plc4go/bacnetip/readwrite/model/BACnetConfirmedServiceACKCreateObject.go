//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
)

// The data-structure of this message
type BACnetConfirmedServiceACKCreateObject struct {
    BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKCreateObject interface {
    IBACnetConfirmedServiceACK
    Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceACKCreateObject) ServiceChoice() uint8 {
    return 0x0A
}

func (m BACnetConfirmedServiceACKCreateObject) initialize() spi.Message {
    return m
}

func NewBACnetConfirmedServiceACKCreateObject() BACnetConfirmedServiceACKInitializer {
    return &BACnetConfirmedServiceACKCreateObject{}
}

func CastIBACnetConfirmedServiceACKCreateObject(structType interface{}) IBACnetConfirmedServiceACKCreateObject {
    castFunc := func(typ interface{}) IBACnetConfirmedServiceACKCreateObject {
        if iBACnetConfirmedServiceACKCreateObject, ok := typ.(IBACnetConfirmedServiceACKCreateObject); ok {
            return iBACnetConfirmedServiceACKCreateObject
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetConfirmedServiceACKCreateObject(structType interface{}) BACnetConfirmedServiceACKCreateObject {
    castFunc := func(typ interface{}) BACnetConfirmedServiceACKCreateObject {
        if sBACnetConfirmedServiceACKCreateObject, ok := typ.(BACnetConfirmedServiceACKCreateObject); ok {
            return sBACnetConfirmedServiceACKCreateObject
        }
        if sBACnetConfirmedServiceACKCreateObject, ok := typ.(*BACnetConfirmedServiceACKCreateObject); ok {
            return *sBACnetConfirmedServiceACKCreateObject
        }
        return BACnetConfirmedServiceACKCreateObject{}
    }
    return castFunc(structType)
}

func (m BACnetConfirmedServiceACKCreateObject) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetConfirmedServiceACK.LengthInBits()

    return lengthInBits
}

func (m BACnetConfirmedServiceACKCreateObject) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKCreateObjectParse(io *spi.ReadBuffer) (BACnetConfirmedServiceACKInitializer, error) {

    // Create the instance
    return NewBACnetConfirmedServiceACKCreateObject(), nil
}

func (m BACnetConfirmedServiceACKCreateObject) Serialize(io spi.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetConfirmedServiceACKSerialize(io, m.BACnetConfirmedServiceACK, CastIBACnetConfirmedServiceACK(m), ser)
}
