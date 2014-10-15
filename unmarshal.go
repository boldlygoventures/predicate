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

func (p *And) UnmarshalJSON(data []byte) error {
	var s Set
	var err error

	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*p = And(s)
	return nil
}

func (p *Or) UnmarshalJSON(data []byte) error {
	var s Set
	var err error

	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	*p = Or(s)
	return nil
}

func (p *Set) UnmarshalJSON(data []byte) error {
	var v interface{}
	var err error

	if v, err = unmarshalJSON(data); err != nil {
		return err
	}

	if _, ok := v.([]interface{}); !ok {
		return fmt.Errorf("predicate: Expected []interface{}, but got %T", v)
	}

	for _, e := range v.([]interface{}) {
		if _, ok := e.(Predicate); !ok {
			return fmt.Errorf("predicate: Expected Predicate, but got %T", e)
		}

		*p = append(*p, e.(Predicate))
	}

	return nil
}

func unmarshalJSON(data []byte) (v interface{}, err error) {
	switch data[0] {
	case '{':
		v, err = unmarshalJSONObject(data)
	case '[':
		v, err = unmarshalJSONArray(data)
	default:
		err = json.Unmarshal(data, &v)
	}

	return
}

func unmarshalJSONObject(data []byte) (v []interface{}, err error) {
	var raw map[string]json.RawMessage

	if err = json.Unmarshal(data, &raw); err != nil {
		return
	}

	v = make([]interface{}, 0)

	for k, d := range raw {
		var p Predicate

		switch k {
		// handle named predicates
		case "and", "or", "xor", "not":
			var s Set

			if err = json.Unmarshal(d, &s); err != nil {
				return
			}

			switch k {
			case "and":
				p = And(s)
			case "or":
				p = Or(s)
			case "xor":
				p = Xor(s)
			case "not":
				p = Not(s...)
			}
		// handle unnamed predicates
		default:
			var x interface{}

			if err = json.Unmarshal(d, &x); err != nil {
				return
			}

			p = Exists(k, x)
		}

		v = append(v, p)
	}

	return
}

func unmarshalJSONArray(data []byte) (v []interface{}, err error) {
	var raw []json.RawMessage

	if err = json.Unmarshal(data, &raw); err != nil {
		return
	}

	v = make([]interface{}, 0)

	for _, d := range raw {
		switch d[0] {
		case '{':
			var s Set

			if err = json.Unmarshal(d, &s); err != nil {
				return
			}

			v = append(v, And(s))
		default:
			var x interface{}

			if x, err = unmarshalJSON(d); err != nil {
				return
			}

			v = append(v, x)
		}
	}

	return
}
