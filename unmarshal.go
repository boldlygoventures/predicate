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

/*
// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *And) UnmarshalJSON(data []byte) error {
	var v Set

	if err := unmarshalJSON(data, &v); err != nil {
		return err
	}

	*p = And(v)
	return nil
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *Or) UnmarshalJSON(data []byte) error {
	var v Set

	if err := unmarshalJSON(data, &v); err != nil {
		return err
	}

	*p = Or(v)
	return nil
}

// UnmarshalJSON satisfies the json.Unmarshaler interface.
func (p *Xor) UnmarshalJSON(data []byte) error {
	var v Set

	if err := unmarshalJSON(data, &v); err != nil {
		return err
	}

	*p = Xor(v)
	return nil
}
*/

func (p *Set) UnmarshalJSON(data []byte) error {
	var v interface{}
	fmt.Printf("> predicate.UnmarshalJSON():\t%v -> %#v\n", string(data), v)

	if err := unmarshalJSON(data, &v); err != nil {
		return err
	}

	*p = make(Set, 0)

	fmt.Printf("< predicate.UnmarshalJSON():\t%v -> %#v\n", string(data), v)
	return nil
}

func unmarshalJSON(data []byte, v interface{}) error {
	fmt.Printf("> unmarshalJSON():\t\t\t\t%v -> %#v\n", string(data), v)
	var err error

	switch data[0] {
	case '{':
		err = unmarshalJSONObject(data, v)
	case '[':
		err = unmarshalJSONArray(data, v)
	default:
		err = json.Unmarshal(data, v)
	}

	fmt.Printf("< unmarshalJSON():\t\t\t\t%v -> %#v\n", string(data), v)
	return err
}

func unmarshalJSONObject(data []byte, v interface{}) error {
	fmt.Printf("> unmarshalJSONObject():\t\t%v -> %#v\n", string(data), v)
	var raw map[string]json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s := make(Set, 0)

	for k, d := range raw {
		var p Predicate

		switch k {
		// handle named predicates
		case "and", "or", "xor", "not":
			var x Set

			if err := json.Unmarshal(d, &x); err != nil {
				return err
			}

			switch k {
			case "and":
				p = And(x)
			case "or":
				p = Or(x)
			case "xor":
				p = Xor(x)
			case "not":
				p = Not(x...)
			}
		// handle unnamed predicates
		default:
			var x interface{}

			if err := json.Unmarshal(d, &x); err != nil {
				return err
			}

			p = Exists(k, x)
		}

		s = append(s, p)
	}

	v = &s
	fmt.Printf("< unmarshalJSONObject():\t\t%v -> %#v\n", string(data), v)
	return nil
}

func unmarshalJSONArray(data []byte, v interface{}) error {
	fmt.Printf("> unmarshalJSONArray():\t\t%v -> %#v\n", string(data), v)
	var raw []json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s := make(Set, 0)

	for _, d := range raw {
		switch d[0] {
		case '{':
			var p And

			if err := unmarshalJSON(d, &p); err != nil {
				return err
			}

			s = append(s, p)
		default:
			var x interface{}
			if err := unmarshalJSON(d, &x); err != nil {
				return err
			}
			v = &x
		}
	}

	if len(s) > 0 {
		v = &s
	}

	fmt.Printf("< unmarshalJSONArray():\t\t%v -> %#v\n", string(data), v)
	return nil
}
