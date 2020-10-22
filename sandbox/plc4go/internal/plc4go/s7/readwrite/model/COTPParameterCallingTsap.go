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
    "errors"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
)

// The data-structure of this message
type COTPParameterCallingTsap struct {
    TsapId uint16
    COTPParameter
}

// The corresponding interface
type ICOTPParameterCallingTsap interface {
    ICOTPParameter
    Serialize(io spi.WriteBuffer) error
}

// Accessors for discriminator values.
func (m COTPParameterCallingTsap) ParameterType() uint8 {
    return 0xC1
}

func (m COTPParameterCallingTsap) initialize() spi.Message {
    return m
}

func NewCOTPParameterCallingTsap(tsapId uint16) COTPParameterInitializer {
    return &COTPParameterCallingTsap{TsapId: tsapId}
}

func CastICOTPParameterCallingTsap(structType interface{}) ICOTPParameterCallingTsap {
    castFunc := func(typ interface{}) ICOTPParameterCallingTsap {
        if iCOTPParameterCallingTsap, ok := typ.(ICOTPParameterCallingTsap); ok {
            return iCOTPParameterCallingTsap
        }
        return nil
    }
    return castFunc(structType)
}

func CastCOTPParameterCallingTsap(structType interface{}) COTPParameterCallingTsap {
    castFunc := func(typ interface{}) COTPParameterCallingTsap {
        if sCOTPParameterCallingTsap, ok := typ.(COTPParameterCallingTsap); ok {
            return sCOTPParameterCallingTsap
        }
        if sCOTPParameterCallingTsap, ok := typ.(*COTPParameterCallingTsap); ok {
            return *sCOTPParameterCallingTsap
        }
        return COTPParameterCallingTsap{}
    }
    return castFunc(structType)
}

func (m COTPParameterCallingTsap) LengthInBits() uint16 {
    var lengthInBits uint16 = m.COTPParameter.LengthInBits()

    // Simple field (tsapId)
    lengthInBits += 16

    return lengthInBits
}

func (m COTPParameterCallingTsap) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPParameterCallingTsapParse(io *spi.ReadBuffer) (COTPParameterInitializer, error) {

    // Simple Field (tsapId)
    tsapId, _tsapIdErr := io.ReadUint16(16)
    if _tsapIdErr != nil {
        return nil, errors.New("Error parsing 'tsapId' field " + _tsapIdErr.Error())
    }

    // Create the instance
    return NewCOTPParameterCallingTsap(tsapId), nil
}

func (m COTPParameterCallingTsap) Serialize(io spi.WriteBuffer) error {
    ser := func() error {

    // Simple Field (tsapId)
    tsapId := uint16(m.TsapId)
    _tsapIdErr := io.WriteUint16(16, (tsapId))
    if _tsapIdErr != nil {
        return errors.New("Error serializing 'tsapId' field " + _tsapIdErr.Error())
    }

        return nil
    }
    return COTPParameterSerialize(io, m.COTPParameter, CastICOTPParameter(m), ser)
}
