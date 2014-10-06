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
	var v []Predicate
	err := unmarshalJSON(data, &v)
	//	*p = And(v.([]Predicate))
	return err
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *Or) UnmarshalJSON(data []byte) error {
	var v []Predicate
	err := unmarshalJSON(data, &v)
	//	*p = Or(v.([]Predicate))
	return err
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *Xor) UnmarshalJSON(data []byte) error {
	var v []Predicate
	err := unmarshalJSON(data, &v)
	//	*p = Xor(v.([]Predicate))
	return err
}

func unmarshalJSON(data []byte, v interface{}) error {
	switch data[0] {
	case '{':
		if err := unmarshalJSONObject(data, v); err != nil {
			return err
		}
		fmt.Printf("object: %#v\n", v)
	case '[':
		if err := unmarshalJSONArray(data, v); err != nil {
			return err
		}
		fmt.Printf("array: %#v\n", v)
	default:
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}

		fmt.Printf("value: %#v\n", v)
	}

	return nil
}

func unmarshalJSONObject(data []byte, v interface{}) error {
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ps := make([]Predicate, 0)

	for k, d := range raw {
		var v interface{}
		if err := unmarshalJSON(d, &v); err != nil {
			return err
		}

		switch k {
		case "and":
			fmt.Println(k)
		case "or":
			fmt.Println(k)
		case "xor":
			fmt.Println(k)
		case "not":
			fmt.Println(k)
		default:
			fmt.Println(k)
		}
	}

	v = &ps
	return nil
}

func unmarshalJSONArray(data []byte, v interface{}) error {
	var raw []json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	ps := make([]Predicate, 0)

	for _, d := range raw {
		var v interface{}
		if err := unmarshalJSON(d, &v); err != nil {
			return err
		}
	}

	v = &ps
	return nil
}
