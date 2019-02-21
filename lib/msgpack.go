package lib

import (
	"bytes"

	"github.com/vmihailenco/msgpack"
)

// Message -
type Message struct {
	Content  string `json:"content"`
	UserId   string `json:"userId"`
	UnixTime int64  `json:"unixTime"`
}

// MarshallMsg -
func MarshallMsg(msg *Message) ([]byte, error) {
	var buf bytes.Buffer
	e := msgpack.NewEncoder(&buf).UseJSONTag(true)
	if err := e.Encode(msg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalMsg -
func UnmarshalMsg(b []byte, msg *Message) error {
	d := msgpack.NewDecoder(bytes.NewReader(b))
	d = d.UseJSONTag(true)
	return d.Decode(msg)
}
