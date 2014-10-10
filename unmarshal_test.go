/*
The MIT License (MIT)

Copyright (c) 2014 Boldly Go Ventures

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package predicate

import (
	"encoding/json"
	"testing"
)

func TestSet_UnmarshalJSON(t *testing.T) {
	var (
		data []byte
		v    Set
	)

	/***** Positive Cases *****/
	data = []byte(`[{"and":[{"x":"a"}]}]`)
	if err := json.Unmarshal(data, &v); err != nil {
		t.Error(err)
	}

	if And(v).P(map[string]interface{}{"x": "a"}) != true {
		t.Fail()
	}

	data = []byte(`[{"or":[{"x":"a"}]}]`)
	if err := json.Unmarshal(data, &v); err != nil {
		t.Error(err)
	}

	if And(v).P(map[string]interface{}{"x": "a"}) != true {
		t.Fail()
	}

	data = []byte(`[{"xor":[{"x":"a"}]}]`)
	if err := json.Unmarshal(data, &v); err != nil {
		t.Error(err)
	}

	if And(v).P(map[string]interface{}{"x": "a"}) != true {
		t.Fail()
	}

	data = []byte(`[{"not":[{"x":"a"}]}]`)
	if err := json.Unmarshal(data, &v); err != nil {
		t.Error(err)
	}

	if And(v).P(map[string]interface{}{"x": "a"}) == true {
		t.Fail()
	}

	/***** Negative Cases *****/
	data = nil
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = make([]byte, 0)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(``)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`"`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`{`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[,]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`{n}`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`{"abc":"xyz",}`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[n]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`{null}`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[abc,]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[null]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[{"and":null}]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[{"or":null}]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[{"xor":null}]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[{"not":null}]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}

	data = []byte(`[{"abc":nil}]`)
	if err := json.Unmarshal(data, &v); err == nil {
		t.Fail()
	}
}
