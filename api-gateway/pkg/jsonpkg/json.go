package jsonpkg

import (
	"bytes"
	"encoding/json"
	"io"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(d []byte, v any) error {
	return json.Unmarshal(d, &v)
}

func Encode(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(&v)
}

func Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(&v)
}

func EncodeWithEncoder(v any) (bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	return buf, err
}
