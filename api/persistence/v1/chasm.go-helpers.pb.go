// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package persistence

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type ChasmNode to the protobuf v3 wire format
func (val *ChasmNode) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmNode from the protobuf v3 wire format
func (val *ChasmNode) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmNode) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmNode values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmNode) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmNode
	switch t := that.(type) {
	case *ChasmNode:
		that1 = t
	case ChasmNode:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmNodeMetadata to the protobuf v3 wire format
func (val *ChasmNodeMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmNodeMetadata from the protobuf v3 wire format
func (val *ChasmNodeMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmNodeMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmNodeMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmNodeMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmNodeMetadata
	switch t := that.(type) {
	case *ChasmNodeMetadata:
		that1 = t
	case ChasmNodeMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmComponentAttributes to the protobuf v3 wire format
func (val *ChasmComponentAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmComponentAttributes from the protobuf v3 wire format
func (val *ChasmComponentAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmComponentAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmComponentAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmComponentAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmComponentAttributes
	switch t := that.(type) {
	case *ChasmComponentAttributes:
		that1 = t
	case ChasmComponentAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmDataAttributes to the protobuf v3 wire format
func (val *ChasmDataAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmDataAttributes from the protobuf v3 wire format
func (val *ChasmDataAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmDataAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmDataAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmDataAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmDataAttributes
	switch t := that.(type) {
	case *ChasmDataAttributes:
		that1 = t
	case ChasmDataAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmCollectionAttributes to the protobuf v3 wire format
func (val *ChasmCollectionAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmCollectionAttributes from the protobuf v3 wire format
func (val *ChasmCollectionAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmCollectionAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmCollectionAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmCollectionAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmCollectionAttributes
	switch t := that.(type) {
	case *ChasmCollectionAttributes:
		that1 = t
	case ChasmCollectionAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmPointerAttributes to the protobuf v3 wire format
func (val *ChasmPointerAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmPointerAttributes from the protobuf v3 wire format
func (val *ChasmPointerAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmPointerAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmPointerAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmPointerAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmPointerAttributes
	switch t := that.(type) {
	case *ChasmPointerAttributes:
		that1 = t
	case ChasmPointerAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmComponentRef to the protobuf v3 wire format
func (val *ChasmComponentRef) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmComponentRef from the protobuf v3 wire format
func (val *ChasmComponentRef) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmComponentRef) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmComponentRef values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmComponentRef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmComponentRef
	switch t := that.(type) {
	case *ChasmComponentRef:
		that1 = t
	case ChasmComponentRef:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChasmTaskInfo to the protobuf v3 wire format
func (val *ChasmTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChasmTaskInfo from the protobuf v3 wire format
func (val *ChasmTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChasmTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChasmTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChasmTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChasmTaskInfo
	switch t := that.(type) {
	case *ChasmTaskInfo:
		that1 = t
	case ChasmTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
