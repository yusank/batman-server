package lib

import (
	"github.com/vmihailenco/msgpack"
)

type Message struct {
	Content  string
	UserId   string
	UnixTime int64
}

func MarshallMsg(msg *Message) ([]byte, error) {
	return msgpack.Marshal(msg)
}

func UnmarshalMsg(b []byte, msg *Message) error {
	return msgpack.Unmarshal(b, msg)
}
