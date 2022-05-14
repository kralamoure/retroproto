package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ChatSubscribeChannelAdd struct {
	Channels []rune
}

func NewChatSubscribeChannelAdd(extra string) (ChatSubscribeChannelAdd, error) {
	var m ChatSubscribeChannelAdd

	if err := m.Deserialize(extra); err != nil {
		return ChatSubscribeChannelAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatSubscribeChannelAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatSubscribeChannelAdd
}

func (m ChatSubscribeChannelAdd) MessageName() string {
	return "ChatSubscribeChannelAdd"
}

func (m ChatSubscribeChannelAdd) Serialized() (string, error) {
	sb := &strings.Builder{}
	for _, channel := range m.Channels {
		sb.WriteRune(channel)
	}
	return sb.String(), nil
}

func (m *ChatSubscribeChannelAdd) Deserialize(extra string) error {
	if extra == "" {
		return nil
	}
	m.Channels = make([]rune, len(extra))
	for i := range extra {
		m.Channels[i] = rune(extra[i])
	}
	return nil
}
