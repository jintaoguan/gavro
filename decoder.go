// Copyright [2019] LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with the
// License.  You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package goavro

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type Decoder struct {
	writerSchema string
	readerSchema string

	// TODO
	// resolver     *resolver

	typeName *name

	nativeFromBinary func([]byte) (interface{}, []byte, error)

	Rabin uint64
}

func NewDecoder(writerSchemaSpecification string, readerSchemaSpecification string) (*Decoder, error) {
	var writerSchema interface{}

	if err := json.Unmarshal([]byte(writerSchemaSpecification), &writerSchema); err != nil {
		return nil, fmt.Errorf("cannot unmarshal schema JSON: %s", err.Error())
	}

	// bootstrap a symbol table with primitive type codecs for the new codec
	st := newSymbolTable()

	c, err := buildCodec(st, nullNamespace, writerSchema)
	if err != nil {
		return nil, err
	}
	c.schemaCanonical, err = parsingCanonicalForm(writerSchema)
	if err != nil {
		return nil, err // should not get here because schema was validated above
	}

	c.Rabin = rabin([]byte(c.schemaCanonical))
	c.soeHeader = []byte{0xC3, 0x01, 0, 0, 0, 0, 0, 0, 0, 0}
	binary.LittleEndian.PutUint64(c.soeHeader[2:], c.Rabin)

	c.schemaOriginal = writerSchemaSpecification
	return nil, nil
}
