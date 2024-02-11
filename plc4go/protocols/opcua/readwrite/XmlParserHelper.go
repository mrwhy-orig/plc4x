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

package readwrite

import (
	"context"
	"strconv"
	"strings"

	"github.com/apache/plc4x/plc4go/protocols/opcua/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type OpcuaXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m OpcuaXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (any, error) {
	switch typeName {
	case "LocaleId":
		return model.LocaleIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ImageGIF":
		return model.ImageGIFParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "EncodedTicket":
		return model.EncodedTicketParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "OpenChannelMessage":
		response := parserArguments[0] == "true"
		return model.OpenChannelMessageParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), response)
	case "ImageJPG":
		return model.ImageJPGParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PascalByteString":
		return model.PascalByteStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DiagnosticInfo":
		return model.DiagnosticInfoParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "PascalString":
		return model.PascalStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TwoByteNodeId":
		return model.TwoByteNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "OpcuaAPU":
		response := parserArguments[0] == "true"
		return model.OpcuaAPUParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), response)
	case "Index":
		return model.IndexParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StatusCode":
		return model.StatusCodeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NormalizedString":
		return model.NormalizedStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "QualifiedName":
		return model.QualifiedNameParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NumericNodeId":
		return model.NumericNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "FourByteNodeId":
		return model.FourByteNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "AudioDataType":
		return model.AudioDataTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SecurityHeader":
		return model.SecurityHeaderParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "UserIdentityTokenDefinition":
		// TODO: find a way to parse the sub types
		var identifier string
		return model.UserIdentityTokenDefinitionParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), identifier)
	case "ContinuationPoint":
		return model.ContinuationPointParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Variant":
		return model.VariantParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Payload":
		extensible := parserArguments[0] == "true"
		parsedUint1, err := strconv.ParseUint(parserArguments[1], 10, 32)
		if err != nil {
			return nil, err
		}
		byteCount := uint32(parsedUint1)
		return model.PayloadParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), extensible, byteCount)
	case "ExtensionObjectEncodingMask":
		return model.ExtensionObjectEncodingMaskParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DurationString":
		return model.DurationStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Structure":
		return model.StructureParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "OpcuaConstants":
		return model.OpcuaConstantsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtensionHeader":
		return model.ExtensionHeaderParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "UtcTime":
		return model.UtcTimeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "MessagePDU":
		response := parserArguments[0] == "true"
		return model.MessagePDUParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), response)
	case "Counter":
		return model.CounterParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SequenceHeader":
		return model.SequenceHeaderParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NodeId":
		return model.NodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "RsaEncryptedSecret":
		return model.RsaEncryptedSecretParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtensionObject":
		includeEncodingMask := parserArguments[0] == "true"
		return model.ExtensionObjectParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), includeEncodingMask)
	case "LocalizedText":
		return model.LocalizedTextParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "IntegerId":
		return model.IntegerIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ByteStringArray":
		return model.ByteStringArrayParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "Handle":
		return model.HandleParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ImagePNG":
		return model.ImagePNGParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "XmlElement":
		return model.XmlElementParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SessionAuthenticationToken":
		return model.SessionAuthenticationTokenParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DataValue":
		return model.DataValueParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "GuidNodeId":
		return model.GuidNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "GuidValue":
		return model.GuidValueParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TrimmedString":
		return model.TrimmedStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ApplicationInstanceCertificate":
		return model.ApplicationInstanceCertificateParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "BitFieldMaskDataType":
		return model.BitFieldMaskDataTypeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ImageBMP":
		return model.ImageBMPParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ExtensionObjectDefinition":
		// TODO: find a way to parse the sub types
		var identifier string
		return model.ExtensionObjectDefinitionParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), identifier)
	case "ExpandedNodeId":
		return model.ExpandedNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "OpcuaProtocolLimits":
		return model.OpcuaProtocolLimitsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NumericRange":
		return model.NumericRangeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "SemanticVersionString":
		return model.SemanticVersionStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "ByteStringNodeId":
		return model.ByteStringNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "TimeString":
		return model.TimeStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "EccEncryptedSecret":
		return model.EccEncryptedSecretParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "StringNodeId":
		return model.StringNodeIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "VersionTime":
		return model.VersionTimeParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "UriString":
		return model.UriStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DecimalString":
		return model.DecimalStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "NodeIdTypeDefinition":
		return model.NodeIdTypeDefinitionParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	case "DateString":
		return model.DateStringParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
