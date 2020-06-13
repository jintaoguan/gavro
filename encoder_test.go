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
	"testing"
)

func TestIntegerEncoder(t *testing.T) {
	testBinaryEncodePass(t, `"int"`, 0, []byte{0})
	testBinaryEncodePass(t, `"int"`, 1, []byte{2})
	testBinaryEncodePass(t, `"int"`, 2, []byte{4})
	testBinaryEncodePass(t, `"int"`, 3, []byte{6})
}

func TestStringEncoder(t *testing.T) {
	testBinaryEncodePass(t, `"string"`, "", []byte{0})
	testBinaryEncodePass(t, `"string"`, "a", []byte{2, 97})
	testBinaryEncodePass(t, `"string"`, "ab", []byte{4, 97, 98})
	testBinaryEncodePass(t, `"string"`, "abc", []byte{6, 97, 98, 99})
}

func TestRecordEncoder(t *testing.T) {
	data := map[string]interface{}{"f1": "a"}
	testBinaryEncodePass(t, `{"name":"r", "type":"record", "fields":[{"name": "f1", "type": "string"}]}`, data, []byte{2, 97})
}
