package elang

import (
	"bytes"
)

// KeyValue ...
type KeyValue map[string]string

// NewKeyValue ...
func NewKeyValue() KeyValue {
	return make(KeyValue)
}

// String ...
func (kv KeyValue) String() string {
	var buffer bytes.Buffer
	for k, v := range kv {
		buffer.WriteRune('[')
		buffer.WriteString(k)
		buffer.WriteRune(':')
		buffer.WriteString(v)
		buffer.WriteRune(']')
	}
	return buffer.String()
}
