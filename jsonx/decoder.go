package jsonx

import (
	"encoding/json"
	"io"
)

type Decoder struct {
	*json.Decoder
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		Decoder: json.NewDecoder(r),
	}
}

func (dec *Decoder) Decode[T any]() (T, error) {
	var t T
	if err := dec.Decoder.Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}
