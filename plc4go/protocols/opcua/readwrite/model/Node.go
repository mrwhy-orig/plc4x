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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Node is the corresponding interface of Node
type Node interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetNodeClass returns NodeClass (property field)
	GetNodeClass() NodeClass
	// GetBrowseName returns BrowseName (property field)
	GetBrowseName() QualifiedName
	// GetDisplayName returns DisplayName (property field)
	GetDisplayName() LocalizedText
	// GetDescription returns Description (property field)
	GetDescription() LocalizedText
	// GetWriteMask returns WriteMask (property field)
	GetWriteMask() uint32
	// GetUserWriteMask returns UserWriteMask (property field)
	GetUserWriteMask() uint32
	// GetNoOfReferences returns NoOfReferences (property field)
	GetNoOfReferences() int32
	// GetReferences returns References (property field)
	GetReferences() []ExtensionObjectDefinition
}

// NodeExactly can be used when we want exactly this type and not a type which fulfills Node.
// This is useful for switch cases.
type NodeExactly interface {
	Node
	isNode() bool
}

// _Node is the data-structure of this message
type _Node struct {
	*_ExtensionObjectDefinition
	NodeId         NodeId
	NodeClass      NodeClass
	BrowseName     QualifiedName
	DisplayName    LocalizedText
	Description    LocalizedText
	WriteMask      uint32
	UserWriteMask  uint32
	NoOfReferences int32
	References     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_Node) GetIdentifier() string {
	return "260"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_Node) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_Node) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Node) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_Node) GetNodeClass() NodeClass {
	return m.NodeClass
}

func (m *_Node) GetBrowseName() QualifiedName {
	return m.BrowseName
}

func (m *_Node) GetDisplayName() LocalizedText {
	return m.DisplayName
}

func (m *_Node) GetDescription() LocalizedText {
	return m.Description
}

func (m *_Node) GetWriteMask() uint32 {
	return m.WriteMask
}

func (m *_Node) GetUserWriteMask() uint32 {
	return m.UserWriteMask
}

func (m *_Node) GetNoOfReferences() int32 {
	return m.NoOfReferences
}

func (m *_Node) GetReferences() []ExtensionObjectDefinition {
	return m.References
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNode factory function for _Node
func NewNode(nodeId NodeId, nodeClass NodeClass, browseName QualifiedName, displayName LocalizedText, description LocalizedText, writeMask uint32, userWriteMask uint32, noOfReferences int32, references []ExtensionObjectDefinition) *_Node {
	_result := &_Node{
		NodeId:                     nodeId,
		NodeClass:                  nodeClass,
		BrowseName:                 browseName,
		DisplayName:                displayName,
		Description:                description,
		WriteMask:                  writeMask,
		UserWriteMask:              userWriteMask,
		NoOfReferences:             noOfReferences,
		References:                 references,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNode(structType any) Node {
	if casted, ok := structType.(Node); ok {
		return casted
	}
	if casted, ok := structType.(*Node); ok {
		return *casted
	}
	return nil
}

func (m *_Node) GetTypeName() string {
	return "Node"
}

func (m *_Node) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (nodeClass)
	lengthInBits += 32

	// Simple field (browseName)
	lengthInBits += m.BrowseName.GetLengthInBits(ctx)

	// Simple field (displayName)
	lengthInBits += m.DisplayName.GetLengthInBits(ctx)

	// Simple field (description)
	lengthInBits += m.Description.GetLengthInBits(ctx)

	// Simple field (writeMask)
	lengthInBits += 32

	// Simple field (userWriteMask)
	lengthInBits += 32

	// Simple field (noOfReferences)
	lengthInBits += 32

	// Array field
	if len(m.References) > 0 {
		for _curItem, element := range m.References {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.References), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_Node) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeParse(ctx context.Context, theBytes []byte, identifier string) (Node, error) {
	return NodeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NodeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (Node, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Node"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Node")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
	_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of Node")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Simple Field (nodeClass)
	if pullErr := readBuffer.PullContext("nodeClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeClass")
	}
	_nodeClass, _nodeClassErr := NodeClassParseWithBuffer(ctx, readBuffer)
	if _nodeClassErr != nil {
		return nil, errors.Wrap(_nodeClassErr, "Error parsing 'nodeClass' field of Node")
	}
	nodeClass := _nodeClass
	if closeErr := readBuffer.CloseContext("nodeClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeClass")
	}

	// Simple Field (browseName)
	if pullErr := readBuffer.PullContext("browseName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for browseName")
	}
	_browseName, _browseNameErr := QualifiedNameParseWithBuffer(ctx, readBuffer)
	if _browseNameErr != nil {
		return nil, errors.Wrap(_browseNameErr, "Error parsing 'browseName' field of Node")
	}
	browseName := _browseName.(QualifiedName)
	if closeErr := readBuffer.CloseContext("browseName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for browseName")
	}

	// Simple Field (displayName)
	if pullErr := readBuffer.PullContext("displayName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for displayName")
	}
	_displayName, _displayNameErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _displayNameErr != nil {
		return nil, errors.Wrap(_displayNameErr, "Error parsing 'displayName' field of Node")
	}
	displayName := _displayName.(LocalizedText)
	if closeErr := readBuffer.CloseContext("displayName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for displayName")
	}

	// Simple Field (description)
	if pullErr := readBuffer.PullContext("description"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for description")
	}
	_description, _descriptionErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _descriptionErr != nil {
		return nil, errors.Wrap(_descriptionErr, "Error parsing 'description' field of Node")
	}
	description := _description.(LocalizedText)
	if closeErr := readBuffer.CloseContext("description"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for description")
	}

	// Simple Field (writeMask)
	_writeMask, _writeMaskErr := readBuffer.ReadUint32("writeMask", 32)
	if _writeMaskErr != nil {
		return nil, errors.Wrap(_writeMaskErr, "Error parsing 'writeMask' field of Node")
	}
	writeMask := _writeMask

	// Simple Field (userWriteMask)
	_userWriteMask, _userWriteMaskErr := readBuffer.ReadUint32("userWriteMask", 32)
	if _userWriteMaskErr != nil {
		return nil, errors.Wrap(_userWriteMaskErr, "Error parsing 'userWriteMask' field of Node")
	}
	userWriteMask := _userWriteMask

	// Simple Field (noOfReferences)
	_noOfReferences, _noOfReferencesErr := readBuffer.ReadInt32("noOfReferences", 32)
	if _noOfReferencesErr != nil {
		return nil, errors.Wrap(_noOfReferencesErr, "Error parsing 'noOfReferences' field of Node")
	}
	noOfReferences := _noOfReferences

	// Array field (references)
	if pullErr := readBuffer.PullContext("references", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for references")
	}
	// Count array
	references := make([]ExtensionObjectDefinition, utils.Max(noOfReferences, 0))
	// This happens when the size is set conditional to 0
	if len(references) == 0 {
		references = nil
	}
	{
		_numItems := uint16(utils.Max(noOfReferences, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "287")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'references' field of Node")
			}
			references[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("references", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for references")
	}

	if closeErr := readBuffer.CloseContext("Node"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Node")
	}

	// Create a partially initialized instance
	_child := &_Node{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NodeId:                     nodeId,
		NodeClass:                  nodeClass,
		BrowseName:                 browseName,
		DisplayName:                displayName,
		Description:                description,
		WriteMask:                  writeMask,
		UserWriteMask:              userWriteMask,
		NoOfReferences:             noOfReferences,
		References:                 references,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_Node) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Node) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("Node"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for Node")
		}

		// Simple Field (nodeId)
		if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeId")
		}
		_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
		if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeId")
		}
		if _nodeIdErr != nil {
			return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
		}

		// Simple Field (nodeClass)
		if pushErr := writeBuffer.PushContext("nodeClass"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeClass")
		}
		_nodeClassErr := writeBuffer.WriteSerializable(ctx, m.GetNodeClass())
		if popErr := writeBuffer.PopContext("nodeClass"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeClass")
		}
		if _nodeClassErr != nil {
			return errors.Wrap(_nodeClassErr, "Error serializing 'nodeClass' field")
		}

		// Simple Field (browseName)
		if pushErr := writeBuffer.PushContext("browseName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for browseName")
		}
		_browseNameErr := writeBuffer.WriteSerializable(ctx, m.GetBrowseName())
		if popErr := writeBuffer.PopContext("browseName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for browseName")
		}
		if _browseNameErr != nil {
			return errors.Wrap(_browseNameErr, "Error serializing 'browseName' field")
		}

		// Simple Field (displayName)
		if pushErr := writeBuffer.PushContext("displayName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for displayName")
		}
		_displayNameErr := writeBuffer.WriteSerializable(ctx, m.GetDisplayName())
		if popErr := writeBuffer.PopContext("displayName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for displayName")
		}
		if _displayNameErr != nil {
			return errors.Wrap(_displayNameErr, "Error serializing 'displayName' field")
		}

		// Simple Field (description)
		if pushErr := writeBuffer.PushContext("description"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for description")
		}
		_descriptionErr := writeBuffer.WriteSerializable(ctx, m.GetDescription())
		if popErr := writeBuffer.PopContext("description"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for description")
		}
		if _descriptionErr != nil {
			return errors.Wrap(_descriptionErr, "Error serializing 'description' field")
		}

		// Simple Field (writeMask)
		writeMask := uint32(m.GetWriteMask())
		_writeMaskErr := writeBuffer.WriteUint32("writeMask", 32, (writeMask))
		if _writeMaskErr != nil {
			return errors.Wrap(_writeMaskErr, "Error serializing 'writeMask' field")
		}

		// Simple Field (userWriteMask)
		userWriteMask := uint32(m.GetUserWriteMask())
		_userWriteMaskErr := writeBuffer.WriteUint32("userWriteMask", 32, (userWriteMask))
		if _userWriteMaskErr != nil {
			return errors.Wrap(_userWriteMaskErr, "Error serializing 'userWriteMask' field")
		}

		// Simple Field (noOfReferences)
		noOfReferences := int32(m.GetNoOfReferences())
		_noOfReferencesErr := writeBuffer.WriteInt32("noOfReferences", 32, (noOfReferences))
		if _noOfReferencesErr != nil {
			return errors.Wrap(_noOfReferencesErr, "Error serializing 'noOfReferences' field")
		}

		// Array Field (references)
		if pushErr := writeBuffer.PushContext("references", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for references")
		}
		for _curItem, _element := range m.GetReferences() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReferences()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'references' field")
			}
		}
		if popErr := writeBuffer.PopContext("references", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for references")
		}

		if popErr := writeBuffer.PopContext("Node"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for Node")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_Node) isNode() bool {
	return true
}

func (m *_Node) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
