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
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type BACnetAddress struct {
	Address []uint8
	Port    uint16
}

// The corresponding interface
type IBACnetAddress interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

func NewBACnetAddress(address []uint8, port uint16) *BACnetAddress {
	return &BACnetAddress{Address: address, Port: port}
}

func CastBACnetAddress(structType interface{}) *BACnetAddress {
	castFunc := func(typ interface{}) *BACnetAddress {
		if casted, ok := typ.(BACnetAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetAddress); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetAddress) GetTypeName() string {
	return "BACnetAddress"
}

func (m *BACnetAddress) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetAddress) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Address) > 0 {
		lengthInBits += 8 * uint16(len(m.Address))
	}

	// Simple field (port)
	lengthInBits += 16

	return lengthInBits
}

func (m *BACnetAddress) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetAddressParse(io utils.ReadBuffer) (*BACnetAddress, error) {
	io.PullContext("BACnetAddress")
	io.PullContext("address")

	// Array field (address)
	// Count array
	address := make([]uint8, uint16(4))
	for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
		_item, _err := io.ReadUint8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'address' field")
		}
		address[curItem] = _item
	}
	io.CloseContext("address")

	// Simple Field (port)
	port, _portErr := io.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field")
	}

	io.CloseContext("BACnetAddress")

	// Create the instance
	return NewBACnetAddress(address, port), nil
}

func (m *BACnetAddress) Serialize(io utils.WriteBuffer) error {
	io.PushContext("BACnetAddress")

	// Array Field (address)
	if m.Address != nil {
		io.PushContext("address")
		for _, _element := range m.Address {
			_elementErr := io.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'address' field")
			}
		}
		io.PopContext("address")
	}

	// Simple Field (port)
	port := uint16(m.Port)
	_portErr := io.WriteUint16("port", 16, (port))
	if _portErr != nil {
		return errors.Wrap(_portErr, "Error serializing 'port' field")
	}

	io.PopContext("BACnetAddress")
	return nil
}

func (m *BACnetAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	for {
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "address":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Address = data
			case "port":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Port = data
			}
		}
	}
}

func (m *BACnetAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.bacnetip.readwrite.BACnetAddress"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Port, xml.StartElement{Name: xml.Name{Local: "port"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m BACnetAddress) String() string {
	return string(m.Box("", 120))
}

func (m BACnetAddress) Box(name string, width int) utils.AsciiBox {
	boxName := "BACnetAddress"
	if name != "" {
		boxName += "/" + name
	}
	boxes := make([]utils.AsciiBox, 0)
	// Array Field (address)
	if m.Address != nil {
		// Simple array base type uint8 will be hex dumped
		boxes = append(boxes, utils.BoxedDumpAnything("Address", m.Address))
		// Simple array base type uint8 will be rendered one by one
		arrayBoxes := make([]utils.AsciiBox, 0)
		for _, _element := range m.Address {
			arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
		}
		boxes = append(boxes, utils.BoxBox("Address", utils.AlignBoxes(arrayBoxes, width-4), 0))
	}
	// Simple field (case simple)
	// uint16 can be boxed as anything with the least amount of space
	boxes = append(boxes, utils.BoxAnything("Port", m.Port, -1))
	return utils.BoxBox(boxName, utils.AlignBoxes(boxes, width-2), 0)
}
