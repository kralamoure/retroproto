package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ChatSubscribeChannelRemove struct {
	Channels []rune
}

func NewChatSubscribeChannelRemove(extra string) (ChatSubscribeChannelRemove, error) {
	var m ChatSubscribeChannelRemove

	if err := m.Deserialize(extra); err != nil {
		return ChatSubscribeChannelRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatSubscribeChannelRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatSubscribeChannelRemove
}

func (m ChatSubscribeChannelRemove) MessageName() string {
	return "ChatSubscribeChannelRemove"
}

func (m ChatSubscribeChannelRemove) Serialized() (string, error) {
	sb := &strings.Builder{}
	for _, channel := range m.Channels {
		sb.WriteRune(channel)
	}
	return sb.String(), nil
}

func (m *ChatSubscribeChannelRemove) Deserialize(extra string) error {
	if extra == "" {
		return nil
	}
	m.Channels = make([]rune, len(extra))
	for i := range extra {
		m.Channels[i] = rune(extra[i])
	}
	return nil
}
