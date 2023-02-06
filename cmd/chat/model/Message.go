package model

import (
	"encoding"
	"encoding/json"
	"time"
)

// Message 聊天记录
type Message struct {
	// Id 记录ID
	Id int64
	// 信息内容
	Content string
	// 发送时间
	CreateTime time.Time
}

// MarshalBinary 实现了 encoding.BinaryMarshaler 接口
func (msg *Message) MarshalBinary() (data []byte, err error) {
	return json.Marshal(msg)
}

var _ encoding.BinaryMarshaler = (*Message)(nil)

// UnmarshalBinary 实现了 encoding.BinaryUnmarshaler 接口
func (msg *Message) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, msg)
}

var _ encoding.BinaryUnmarshaler = (*Message)(nil)
