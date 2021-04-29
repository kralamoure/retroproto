package msgcli

import (
	"strings"

	"github.com/kralamoure/d1encoding"
)

type ChatRequestSubscribeChannelAdd struct {
	Channels []rune
}

func (m ChatRequestSubscribeChannelAdd) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ChatRequestSubscribeChannelAdd
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
