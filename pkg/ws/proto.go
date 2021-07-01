package ws

import (
	"fmt"
	"strings"

	"github.com/arthur-snake/snakego/pkg/proto"
)

func convertCellUpdate(c proto.UpdateCell) string {
	str := fmt.Sprintf("%d.%d#%s", c.Location.X, c.Location.Y, c.ID)
	if c.Food > 0 { // TODO: omitempty?
		str += fmt.Sprintf("*%d", c.Food)
	}
	return str
}

func convertCellUpdates(upds []proto.UpdateCell) string {
	var m []string
	for _, u := range upds {
		m = append(m, convertCellUpdate(u))
	}

	return strings.Join(m, "|")
}

type serverMessage struct {
	Act  string       `json:"act"`
	Map  string       `json:"a,omitempty"`
	ID   []idInfo     `json:"u,omitempty"`
	Chat []chatUpdate `json:"c,omitempty"`

	SizeX int `json:"columns"`
	SizeY int `json:"rows"`
}

func convertInitMessage(message proto.InitMessage) serverMessage {
	msg := convertUpdateMessage(*message.Update)
	msg.Act = "init"
	msg.SizeX = message.Size.SizeX
	msg.SizeY = message.Size.SizeY

	return msg
}

func convertUpdateMessage(message proto.UpdateMessage) serverMessage {
	return serverMessage{
		Act:  "upd",
		Map:  convertCellUpdates(message.CellUpdates),
		ID:   convertIDUpdates(message.IDUpdates),
		Chat: convertChatUpdates(message.ChatUpdates),
	}
}

type idInfo struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Nick  string `json:"nick"`
}

func convertIDUpdate(id proto.UpdateID) idInfo {
	return idInfo{
		ID:    string(id.ID),
		Type:  string(id.Type),
		Color: string(id.Color),
		Nick:  id.Nick,
	}
}

func convertIDUpdates(arr []proto.UpdateID) []idInfo {
	var res []idInfo
	for _, upd := range arr {
		res = append(res, convertIDUpdate(upd))
	}
	return res
}

type chatUpdate struct {
	ID      string `json:"id"`
	Message string `json:"msg"`
}

func convertChatUpdate(chat proto.UpdateChat) chatUpdate {
	return chatUpdate{
		ID:      string(chat.ID),
		Message: chat.Message,
	}
}

func convertChatUpdates(arr []proto.UpdateChat) []chatUpdate {
	var res []chatUpdate
	for _, upd := range arr {
		res = append(res, convertChatUpdate(upd))
	}
	return res
}
