package msgcli

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ChatRequestSubscribeChannelAdd struct {
	Channels []rune
}

func NewChatRequestSubscribeChannelAdd(extra string) (ChatRequestSubscribeChannelAdd, error) {
	var m ChatRequestSubscribeChannelAdd

	if err := m.Deserialize(extra); err != nil {
		return ChatRequestSubscribeChannelAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatRequestSubscribeChannelAdd) MessageId() retroproto.MsgCliId {
	return retroproto.ChatRequestSubscribeChannelAdd
}

func (m ChatRequestSubscribeChannelAdd) MessageName() string {
	return "ChatRequestSubscribeChannelAdd"
}

func (m ChatRequestSubscribeChannelAdd) Serialized() (string, error) {
	sb := &strings.Builder{}
	for _, channel := range m.Channels {
		sb.WriteRune(channel)
	}
	return sb.String(), nil
}

func (m *ChatRequestSubscribeChannelAdd) Deserialize(extra string) error {
	if extra == "" {
		return nil
	}
	m.Channels = make([]rune, len(extra))
	for i := range extra {
		m.Channels[i] = rune(extra[i])
	}
	return nil
}
