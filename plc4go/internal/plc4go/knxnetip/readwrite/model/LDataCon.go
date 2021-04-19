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
type LDataCon struct {
	AdditionalInformationLength uint8
	AdditionalInformation       []*CEMIAdditionalInformation
	DataFrame                   *LDataFrame
	Parent                      *CEMI
}

// The corresponding interface
type ILDataCon interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *LDataCon) MessageCode() uint8 {
	return 0x2E
}

func (m *LDataCon) InitializeParent(parent *CEMI) {
}

func NewLDataCon(additionalInformationLength uint8, additionalInformation []*CEMIAdditionalInformation, dataFrame *LDataFrame) *CEMI {
	child := &LDataCon{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Parent:                      NewCEMI(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastLDataCon(structType interface{}) *LDataCon {
	castFunc := func(typ interface{}) *LDataCon {
		if casted, ok := typ.(LDataCon); ok {
			return &casted
		}
		if casted, ok := typ.(*LDataCon); ok {
			return casted
		}
		if casted, ok := typ.(CEMI); ok {
			return CastLDataCon(casted.Child)
		}
		if casted, ok := typ.(*CEMI); ok {
			return CastLDataCon(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *LDataCon) GetTypeName() string {
	return "LDataCon"
}

func (m *LDataCon) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *LDataCon) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.LengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.LengthInBits()

	return lengthInBits
}

func (m *LDataCon) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func LDataConParse(io utils.ReadBuffer) (*CEMI, error) {
	io.PullContext("LDataCon")

	// Simple Field (additionalInformationLength)
	additionalInformationLength, _additionalInformationLengthErr := io.ReadUint8("additionalInformationLength", 8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field")
	}
	io.PullContext("additionalInformation")

	// Array field (additionalInformation)
	// Length array
	additionalInformation := make([]*CEMIAdditionalInformation, 0)
	_additionalInformationLength := additionalInformationLength
	_additionalInformationEndPos := io.GetPos() + uint16(_additionalInformationLength)
	for io.GetPos() < _additionalInformationEndPos {
		_item, _err := CEMIAdditionalInformationParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field")
		}
		additionalInformation = append(additionalInformation, _item)
	}
	io.CloseContext("additionalInformation")

	// Simple Field (dataFrame)
	dataFrame, _dataFrameErr := LDataFrameParse(io)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field")
	}

	io.CloseContext("LDataCon")

	// Create a partially initialized instance
	_child := &LDataCon{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation:       additionalInformation,
		DataFrame:                   dataFrame,
		Parent:                      &CEMI{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *LDataCon) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		io.PushContext("LDataCon")

		// Simple Field (additionalInformationLength)
		additionalInformationLength := uint8(m.AdditionalInformationLength)
		_additionalInformationLengthErr := io.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
		if _additionalInformationLengthErr != nil {
			return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
		}

		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			io.PushContext("additionalInformation")
			for _, _element := range m.AdditionalInformation {
				_elementErr := _element.Serialize(io)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
				}
			}
			io.PopContext("additionalInformation")
		}

		// Simple Field (dataFrame)
		_dataFrameErr := m.DataFrame.Serialize(io)
		if _dataFrameErr != nil {
			return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
		}

		io.PopContext("LDataCon")
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *LDataCon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "additionalInformationLength":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.AdditionalInformationLength = data
			case "additionalInformation":
			arrayLoop:
				for {
					token, err = d.Token()
					switch token.(type) {
					case xml.StartElement:
						tok := token.(xml.StartElement)
						var dt *CEMIAdditionalInformation
						if err := d.DecodeElement(&dt, &tok); err != nil {
							return err
						}
						m.AdditionalInformation = append(m.AdditionalInformation, dt)
					default:
						break arrayLoop
					}
				}
			case "dataFrame":
				var dt *LDataFrame
				if err := d.DecodeElement(&dt, &tok); err != nil {
					if err == io.EOF {
						continue
					}
					return err
				}
				m.DataFrame = dt
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *LDataCon) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.AdditionalInformationLength, xml.StartElement{Name: xml.Name{Local: "additionalInformationLength"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	for _, arrayElement := range m.AdditionalInformation {
		if err := e.EncodeElement(arrayElement, xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.DataFrame, xml.StartElement{Name: xml.Name{Local: "dataFrame"}}); err != nil {
		return err
	}
	return nil
}

func (m LDataCon) String() string {
	return string(m.Box("", 120))
}

func (m LDataCon) Box(name string, width int) utils.AsciiBox {
	boxName := "LDataCon"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("AdditionalInformationLength", m.AdditionalInformationLength, -1))
		// Array Field (additionalInformation)
		if m.AdditionalInformation != nil {
			// Complex array base type
			arrayBoxes := make([]utils.AsciiBox, 0)
			for _, _element := range m.AdditionalInformation {
				arrayBoxes = append(arrayBoxes, utils.BoxAnything("", _element, width-2))
			}
			boxes = append(boxes, utils.BoxBox("AdditionalInformation", utils.AlignBoxes(arrayBoxes, width-4), 0))
		}
		// Complex field (case complex)
		boxes = append(boxes, m.DataFrame.Box("dataFrame", width-2))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}
