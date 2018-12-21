package message

import "encoding/xml"

// receive
type Message struct {
	ToUserName   string   `xml:"ToUserName,CDATA"`
	FromUserName string   `xml:"FromUserName,CDATA"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType,CDATA"`
	Content      string   `xml:"Content,CDATA,omitempty"`
	MsgId        string   `xml:"MsgId,CDATA,omitempty"`
	MediaId      string   `xml:"MediaId,omitempty,CDATA"`
}

// reply
type Text struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName,CDATA"`
	FromUserName string   `xml:"FromUserName,CDATA"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType,CDATA"`
	Content      string   `xml:"Content,CDATA,omitempty"`
}

type Image struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName,CDATA"`
	FromUserName string   `xml:"FromUserName,CDATA"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType,CDATA"`
	MediaId      string   `xml:"MediaId,omitempty,CDATA"`
}
