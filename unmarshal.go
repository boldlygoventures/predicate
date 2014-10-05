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
	"fmt"
)

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *And) UnmarshalJSON(data []byte) error {
	v, err := unmarshalJSON(data)
	*p = And(v.([]Predicate))
	return err
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *Or) UnmarshalJSON(data []byte) error {
	v, err := unmarshalJSON(data)
	*p = Or(v.([]Predicate))
	return err
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *Xor) UnmarshalJSON(data []byte) error {
	v, err := unmarshalJSON(data)
	*p = Xor(v.([]Predicate))
	return err
}

func unmarshalJSON(data []byte) (interface{}, error) {
	switch data[0] {
	case '{':
		fmt.Println("object")
		if err := unmarshalJSONObject(data); err != nil {
			return nil, err
		}
	case '[':
		fmt.Println("array")
		if err := unmarshalJSONArray(data); err != nil {
			return nil, err
		}
	default:
		var v interface{}
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}

		fmt.Println("value", v)
	}

	return nil, nil
}

func unmarshalJSONObject(data []byte) error {
	var v map[string]json.RawMessage

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	for k, v := range v {
		fmt.Println(k)
		if _, err := unmarshalJSON(v); err != nil {
			return err
		}
	}
	return nil
}

func unmarshalJSONArray(data []byte) error {
	var v []json.RawMessage

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	for _, v := range v {
		if _, err := unmarshalJSON(v); err != nil {
			return err
		}
	}

	return nil
}
