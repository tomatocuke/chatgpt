package wechat

import (
	"encoding/xml"
	"sync"
	"time"
)

var (
	cache sync.Map
)

type Msg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`

	MsgId int64 `xml:"MsgId,omitempty"`
}

func NewMsg(data []byte) *Msg {
	var msg Msg
	if err := xml.Unmarshal(data, &msg); err != nil {
		return nil
	}
	return &msg
}

func (msg *Msg) IsText() bool {
	return msg.MsgType == "text"
}

func (msg *Msg) HasReply() bool {
	_, ok := cache.Load(msg.FromUserName)
	return ok
}

func (msg *Msg) GenerateEchoData(s string) []byte {
	data := Msg{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      msg.MsgType,
		Content:      s,
	}
	bs, _ := xml.Marshal(&data)
	return bs
}
