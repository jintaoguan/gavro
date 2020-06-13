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
	"fmt"
)

type Encoder struct {
	Codec *Codec
}

func NewEncoder(schemaSpecification string) (*Encoder, error) {
	c, err := NewCodec(schemaSpecification)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize codec. %s", err.Error())
	}
	return &Encoder{Codec: c}, nil
}

func (e *Encoder) Encode(buf []byte, datum interface{}) ([]byte, error) {
	buf, err := e.Codec.binaryFromNative(buf, datum)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize codec. %s", err.Error())
	}
	return buf, nil
}
