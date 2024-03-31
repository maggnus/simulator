package api

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Sender      string `json:"sender"`
	Instruction string `json:"instruction"`
}

func NewMessage(sender string, instruction string) *Message {
	return &Message{Sender: sender, Instruction: instruction}
}

func NewMessageData(b []byte) (*Message, error) {
	var message Message
	err := json.Unmarshal(b, &message)
	if err != nil {
		return nil, message.Valid()
	}

	return &message, nil
}

func (c *Message) Valid() error {
	if c.Sender == "" {
		return fmt.Errorf("message sender is not defined: %+v", c)
	}

	if c.Instruction == "" {
		return fmt.Errorf("message instruction is not defined: %+v", c)
	}

	return nil
}
