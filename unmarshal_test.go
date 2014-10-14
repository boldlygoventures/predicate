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

var (
	pj = [][]byte{
		[]byte(`[{"and":[{"x":"a"}]}]`),
		[]byte(`[{"or":[{"x":"a"}]}]`),
		[]byte(`[{"xor":[{"x":"a"}]}]`),
		[]byte(`[{"not":[{"x":"a"}]}]`),
	}

	nj = [][]byte{
		[]byte(nil),
		make([]byte, 0),
		[]byte(``),
		[]byte(`"`),
		[]byte(`{`),
		[]byte(`[`),
		[]byte(`[,]`),
		[]byte(`{n}`),
		[]byte(`{"abc":"xyz",}`),
		[]byte(`[n]`),
		[]byte(`{null}`),
		[]byte(`[abc,]`),
		[]byte(`[null]`),
		[]byte(`[{"and":null}]`),
		[]byte(`[{"or":null}]`),
		[]byte(`[{"xor":null}]`),
		[]byte(`[{"not":null}]`),
		[]byte(`[{"abc":nil}]`),
	}
)

func TestAnd_UnmarshalJSON(t *testing.T) {
	var v And

	/***** Positive Cases *****/
	for _, data := range pj {
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			t.Error(err)
		}
	}

	/***** Negative Cases *****/
	for _, data := range nj {
		if err := json.Unmarshal(data, &v); err == nil {
			t.Error(err)
		}
	}
}

func TestSet_UnmarshalJSON(t *testing.T) {
	var v Set

	/***** Positive Cases *****/
	for _, data := range pj {
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			t.Error(err)
		}
	}

	/***** Negative Cases *****/
	for _, data := range nj {
		if err := json.Unmarshal(data, &v); err == nil {
			t.Error(err)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var (
		data []byte
	)

	data = []byte(`[{"abc":nil}]`)
	if _, err := unmarshalJSON(data); err == nil {
		t.Fail()
	}
}

func TestUnmarshalJSONObject(t *testing.T) {
	var (
		data []byte
	)

	data = []byte(`[{"abc":nil}]`)
	if _, err := unmarshalJSONObject(data); err == nil {
		t.Fail()
	}

	data = []byte(`{"abc":["a","b","c",]}`)
	if _, err := unmarshalJSONObject(data); err == nil {
		t.Fail()
	}
}

func TestUnmarshalJSONArray(t *testing.T) {
	var (
		data []byte
	)

	data = []byte(`[{"abc":nil}]`)
	if _, err := unmarshalJSONObject(data); err == nil {
		t.Fail()
	}
}
