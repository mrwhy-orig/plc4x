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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
    "strconv"
)

// Constant values.
const BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER uint8 = 0x09
const BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER uint8 = 0x1C
const BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER uint8 = 0x29
const BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS uint8 = 0x00
const BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER uint8 = 0x07

// The data-structure of this message
type BACnetConfirmedServiceRequestSubscribeCOV struct {
    SubscriberProcessIdentifier uint8
    MonitoredObjectType uint16
    MonitoredObjectInstanceNumber uint32
    IssueConfirmedNotifications bool
    LifetimeLength uint8
    LifetimeSeconds []int8
    BACnetConfirmedServiceRequest
}

// The corresponding interface
type IBACnetConfirmedServiceRequestSubscribeCOV interface {
    IBACnetConfirmedServiceRequest
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceRequestSubscribeCOV) ServiceChoice() uint8 {
    return 0x05
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) initialize() spi.Message {
    return m
}

func NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier uint8, monitoredObjectType uint16, monitoredObjectInstanceNumber uint32, issueConfirmedNotifications bool, lifetimeLength uint8, lifetimeSeconds []int8) BACnetConfirmedServiceRequestInitializer {
    return &BACnetConfirmedServiceRequestSubscribeCOV{SubscriberProcessIdentifier: subscriberProcessIdentifier, MonitoredObjectType: monitoredObjectType, MonitoredObjectInstanceNumber: monitoredObjectInstanceNumber, IssueConfirmedNotifications: issueConfirmedNotifications, LifetimeLength: lifetimeLength, LifetimeSeconds: lifetimeSeconds}
}

func CastIBACnetConfirmedServiceRequestSubscribeCOV(structType interface{}) IBACnetConfirmedServiceRequestSubscribeCOV {
    castFunc := func(typ interface{}) IBACnetConfirmedServiceRequestSubscribeCOV {
        if iBACnetConfirmedServiceRequestSubscribeCOV, ok := typ.(IBACnetConfirmedServiceRequestSubscribeCOV); ok {
            return iBACnetConfirmedServiceRequestSubscribeCOV
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetConfirmedServiceRequestSubscribeCOV(structType interface{}) BACnetConfirmedServiceRequestSubscribeCOV {
    castFunc := func(typ interface{}) BACnetConfirmedServiceRequestSubscribeCOV {
        if sBACnetConfirmedServiceRequestSubscribeCOV, ok := typ.(BACnetConfirmedServiceRequestSubscribeCOV); ok {
            return sBACnetConfirmedServiceRequestSubscribeCOV
        }
        if sBACnetConfirmedServiceRequestSubscribeCOV, ok := typ.(*BACnetConfirmedServiceRequestSubscribeCOV); ok {
            return *sBACnetConfirmedServiceRequestSubscribeCOV
        }
        return BACnetConfirmedServiceRequestSubscribeCOV{}
    }
    return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetConfirmedServiceRequest.LengthInBits()

    // Const Field (subscriberProcessIdentifierHeader)
    lengthInBits += 8

    // Simple field (subscriberProcessIdentifier)
    lengthInBits += 8

    // Const Field (monitoredObjectIdentifierHeader)
    lengthInBits += 8

    // Simple field (monitoredObjectType)
    lengthInBits += 10

    // Simple field (monitoredObjectInstanceNumber)
    lengthInBits += 22

    // Const Field (issueConfirmedNotificationsHeader)
    lengthInBits += 8

    // Const Field (issueConfirmedNotificationsSkipBits)
    lengthInBits += 7

    // Simple field (issueConfirmedNotifications)
    lengthInBits += 1

    // Const Field (lifetimeHeader)
    lengthInBits += 5

    // Simple field (lifetimeLength)
    lengthInBits += 3

    // Array field
    if len(m.LifetimeSeconds) > 0 {
        lengthInBits += 8 * uint16(len(m.LifetimeSeconds))
    }

    return lengthInBits
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVParse(io *utils.ReadBuffer) (BACnetConfirmedServiceRequestInitializer, error) {

    // Const Field (subscriberProcessIdentifierHeader)
    subscriberProcessIdentifierHeader, _subscriberProcessIdentifierHeaderErr := io.ReadUint8(8)
    if _subscriberProcessIdentifierHeaderErr != nil {
        return nil, errors.New("Error parsing 'subscriberProcessIdentifierHeader' field " + _subscriberProcessIdentifierHeaderErr.Error())
    }
    if subscriberProcessIdentifierHeader != BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_SUBSCRIBERPROCESSIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(subscriberProcessIdentifierHeader)))
    }

    // Simple Field (subscriberProcessIdentifier)
    subscriberProcessIdentifier, _subscriberProcessIdentifierErr := io.ReadUint8(8)
    if _subscriberProcessIdentifierErr != nil {
        return nil, errors.New("Error parsing 'subscriberProcessIdentifier' field " + _subscriberProcessIdentifierErr.Error())
    }

    // Const Field (monitoredObjectIdentifierHeader)
    monitoredObjectIdentifierHeader, _monitoredObjectIdentifierHeaderErr := io.ReadUint8(8)
    if _monitoredObjectIdentifierHeaderErr != nil {
        return nil, errors.New("Error parsing 'monitoredObjectIdentifierHeader' field " + _monitoredObjectIdentifierHeaderErr.Error())
    }
    if monitoredObjectIdentifierHeader != BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_MONITOREDOBJECTIDENTIFIERHEADER)) + " but got " + strconv.Itoa(int(monitoredObjectIdentifierHeader)))
    }

    // Simple Field (monitoredObjectType)
    monitoredObjectType, _monitoredObjectTypeErr := io.ReadUint16(10)
    if _monitoredObjectTypeErr != nil {
        return nil, errors.New("Error parsing 'monitoredObjectType' field " + _monitoredObjectTypeErr.Error())
    }

    // Simple Field (monitoredObjectInstanceNumber)
    monitoredObjectInstanceNumber, _monitoredObjectInstanceNumberErr := io.ReadUint32(22)
    if _monitoredObjectInstanceNumberErr != nil {
        return nil, errors.New("Error parsing 'monitoredObjectInstanceNumber' field " + _monitoredObjectInstanceNumberErr.Error())
    }

    // Const Field (issueConfirmedNotificationsHeader)
    issueConfirmedNotificationsHeader, _issueConfirmedNotificationsHeaderErr := io.ReadUint8(8)
    if _issueConfirmedNotificationsHeaderErr != nil {
        return nil, errors.New("Error parsing 'issueConfirmedNotificationsHeader' field " + _issueConfirmedNotificationsHeaderErr.Error())
    }
    if issueConfirmedNotificationsHeader != BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSHEADER)) + " but got " + strconv.Itoa(int(issueConfirmedNotificationsHeader)))
    }

    // Const Field (issueConfirmedNotificationsSkipBits)
    issueConfirmedNotificationsSkipBits, _issueConfirmedNotificationsSkipBitsErr := io.ReadUint8(7)
    if _issueConfirmedNotificationsSkipBitsErr != nil {
        return nil, errors.New("Error parsing 'issueConfirmedNotificationsSkipBits' field " + _issueConfirmedNotificationsSkipBitsErr.Error())
    }
    if issueConfirmedNotificationsSkipBits != BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_ISSUECONFIRMEDNOTIFICATIONSSKIPBITS)) + " but got " + strconv.Itoa(int(issueConfirmedNotificationsSkipBits)))
    }

    // Simple Field (issueConfirmedNotifications)
    issueConfirmedNotifications, _issueConfirmedNotificationsErr := io.ReadBit()
    if _issueConfirmedNotificationsErr != nil {
        return nil, errors.New("Error parsing 'issueConfirmedNotifications' field " + _issueConfirmedNotificationsErr.Error())
    }

    // Const Field (lifetimeHeader)
    lifetimeHeader, _lifetimeHeaderErr := io.ReadUint8(5)
    if _lifetimeHeaderErr != nil {
        return nil, errors.New("Error parsing 'lifetimeHeader' field " + _lifetimeHeaderErr.Error())
    }
    if lifetimeHeader != BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BACnetConfirmedServiceRequestSubscribeCOV_LIFETIMEHEADER)) + " but got " + strconv.Itoa(int(lifetimeHeader)))
    }

    // Simple Field (lifetimeLength)
    lifetimeLength, _lifetimeLengthErr := io.ReadUint8(3)
    if _lifetimeLengthErr != nil {
        return nil, errors.New("Error parsing 'lifetimeLength' field " + _lifetimeLengthErr.Error())
    }

    // Array field (lifetimeSeconds)
    // Count array
    lifetimeSeconds := make([]int8, lifetimeLength)
    for curItem := uint16(0); curItem < uint16(lifetimeLength); curItem++ {

        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'lifetimeSeconds' field " + _err.Error())
        }
        lifetimeSeconds[curItem] = _item
    }

    // Create the instance
    return NewBACnetConfirmedServiceRequestSubscribeCOV(subscriberProcessIdentifier, monitoredObjectType, monitoredObjectInstanceNumber, issueConfirmedNotifications, lifetimeLength, lifetimeSeconds), nil
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Const Field (subscriberProcessIdentifierHeader)
    _subscriberProcessIdentifierHeaderErr := io.WriteUint8(8, 0x09)
    if _subscriberProcessIdentifierHeaderErr != nil {
        return errors.New("Error serializing 'subscriberProcessIdentifierHeader' field " + _subscriberProcessIdentifierHeaderErr.Error())
    }

    // Simple Field (subscriberProcessIdentifier)
    subscriberProcessIdentifier := uint8(m.SubscriberProcessIdentifier)
    _subscriberProcessIdentifierErr := io.WriteUint8(8, (subscriberProcessIdentifier))
    if _subscriberProcessIdentifierErr != nil {
        return errors.New("Error serializing 'subscriberProcessIdentifier' field " + _subscriberProcessIdentifierErr.Error())
    }

    // Const Field (monitoredObjectIdentifierHeader)
    _monitoredObjectIdentifierHeaderErr := io.WriteUint8(8, 0x1C)
    if _monitoredObjectIdentifierHeaderErr != nil {
        return errors.New("Error serializing 'monitoredObjectIdentifierHeader' field " + _monitoredObjectIdentifierHeaderErr.Error())
    }

    // Simple Field (monitoredObjectType)
    monitoredObjectType := uint16(m.MonitoredObjectType)
    _monitoredObjectTypeErr := io.WriteUint16(10, (monitoredObjectType))
    if _monitoredObjectTypeErr != nil {
        return errors.New("Error serializing 'monitoredObjectType' field " + _monitoredObjectTypeErr.Error())
    }

    // Simple Field (monitoredObjectInstanceNumber)
    monitoredObjectInstanceNumber := uint32(m.MonitoredObjectInstanceNumber)
    _monitoredObjectInstanceNumberErr := io.WriteUint32(22, (monitoredObjectInstanceNumber))
    if _monitoredObjectInstanceNumberErr != nil {
        return errors.New("Error serializing 'monitoredObjectInstanceNumber' field " + _monitoredObjectInstanceNumberErr.Error())
    }

    // Const Field (issueConfirmedNotificationsHeader)
    _issueConfirmedNotificationsHeaderErr := io.WriteUint8(8, 0x29)
    if _issueConfirmedNotificationsHeaderErr != nil {
        return errors.New("Error serializing 'issueConfirmedNotificationsHeader' field " + _issueConfirmedNotificationsHeaderErr.Error())
    }

    // Const Field (issueConfirmedNotificationsSkipBits)
    _issueConfirmedNotificationsSkipBitsErr := io.WriteUint8(7, 0x00)
    if _issueConfirmedNotificationsSkipBitsErr != nil {
        return errors.New("Error serializing 'issueConfirmedNotificationsSkipBits' field " + _issueConfirmedNotificationsSkipBitsErr.Error())
    }

    // Simple Field (issueConfirmedNotifications)
    issueConfirmedNotifications := bool(m.IssueConfirmedNotifications)
    _issueConfirmedNotificationsErr := io.WriteBit((issueConfirmedNotifications))
    if _issueConfirmedNotificationsErr != nil {
        return errors.New("Error serializing 'issueConfirmedNotifications' field " + _issueConfirmedNotificationsErr.Error())
    }

    // Const Field (lifetimeHeader)
    _lifetimeHeaderErr := io.WriteUint8(5, 0x07)
    if _lifetimeHeaderErr != nil {
        return errors.New("Error serializing 'lifetimeHeader' field " + _lifetimeHeaderErr.Error())
    }

    // Simple Field (lifetimeLength)
    lifetimeLength := uint8(m.LifetimeLength)
    _lifetimeLengthErr := io.WriteUint8(3, (lifetimeLength))
    if _lifetimeLengthErr != nil {
        return errors.New("Error serializing 'lifetimeLength' field " + _lifetimeLengthErr.Error())
    }

    // Array Field (lifetimeSeconds)
    if m.LifetimeSeconds != nil {
        for _, _element := range m.LifetimeSeconds {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'lifetimeSeconds' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return BACnetConfirmedServiceRequestSerialize(io, m.BACnetConfirmedServiceRequest, CastIBACnetConfirmedServiceRequest(m), ser)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "subscriberProcessIdentifier":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SubscriberProcessIdentifier = data
            case "monitoredObjectType":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.MonitoredObjectType = data
            case "monitoredObjectInstanceNumber":
                var data uint32
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.MonitoredObjectInstanceNumber = data
            case "issueConfirmedNotifications":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.IssueConfirmedNotifications = data
            case "lifetimeLength":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.LifetimeLength = data
            case "lifetimeSeconds":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.LifetimeSeconds = utils.ByteToInt8(_decoded[0:_len])
            }
        }
    }
}

func (m BACnetConfirmedServiceRequestSubscribeCOV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceRequestSubscribeCOV"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SubscriberProcessIdentifier, xml.StartElement{Name: xml.Name{Local: "subscriberProcessIdentifier"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.MonitoredObjectType, xml.StartElement{Name: xml.Name{Local: "monitoredObjectType"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.MonitoredObjectInstanceNumber, xml.StartElement{Name: xml.Name{Local: "monitoredObjectInstanceNumber"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.IssueConfirmedNotifications, xml.StartElement{Name: xml.Name{Local: "issueConfirmedNotifications"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.LifetimeLength, xml.StartElement{Name: xml.Name{Local: "lifetimeLength"}}); err != nil {
        return err
    }
    _encodedLifetimeSeconds := make([]byte, base64.StdEncoding.EncodedLen(len(m.LifetimeSeconds)))
    base64.StdEncoding.Encode(_encodedLifetimeSeconds, utils.Int8ToByte(m.LifetimeSeconds))
    if err := e.EncodeElement(_encodedLifetimeSeconds, xml.StartElement{Name: xml.Name{Local: "lifetimeSeconds"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

