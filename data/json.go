package data

import (
	"encoding/json"
	"io"
)

// FromJSON decodes json from an ioReader into an interface
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

// ToJSON returns a JSON representation of an interface
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}