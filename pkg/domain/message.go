package domain

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type CellType string

var (
	FreeCell   CellType = "free"
	FoodCell   CellType = "food"
	PlayerCell CellType = "player"
	BlockCell  CellType = "block"
)

type IDInfo struct {
	ID    string   `json:"id"`
	Type  CellType `json:"type"`
	Color Color    `json:"color"`
	Nick  string   `json:"nick"`
}

type ChatUpdate struct {
}

type Message interface {
	Marshal() []byte
}

type CellUpdate struct {
	X    int
	Y    int
	ID   string
	Food int // TODO: zero food and no food
}

func (c CellUpdate) Marshal() string {
	str := fmt.Sprintf("%d.%d#%s", c.X, c.Y, c.ID)
	if c.Food > 0 {
		str += fmt.Sprintf("*%d", c.Food)
	}
	return str
}

func MarshalMapUpdates(upds []CellUpdate) string {
	var m []string
	for _, u := range upds {
		m = append(m, u.Marshal())
	}

	return strings.Join(m, "|")
}

type InitMessage struct {
	Act         string       `json:"act"` // "init"
	MapUpdates  string       `json:"a,omitempty"`
	IDUpdates   []IDInfo     `json:"u,omitempty"`
	ChatUpdates []ChatUpdate `json:"c,omitempty"`

	SizeX int `json:"columns"`
	SizeY int `json:"rows"`
}

func (i InitMessage) Marshal() []byte {
	res, err := json.Marshal(i)
	if err != nil {
		log.WithError(err).Error("failed to marshal InitMessage")
	}

	return res
}
