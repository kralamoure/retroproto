package msgcli

import (
	"strings"

	"github.com/kralamoure/d1encoding"
)

type ChatRequestSubscribeChannelRemove struct {
	Channels []rune
}

func (m ChatRequestSubscribeChannelRemove) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ChatRequestSubscribeChannelRemove
}

func (m ChatRequestSubscribeChannelRemove) Serialized() (string, error) {
	sb := &strings.Builder{}
	for _, channel := range m.Channels {
		sb.WriteRune(channel)
	}
	return sb.String(), nil
}

func (m *ChatRequestSubscribeChannelRemove) Deserialize(extra string) error {
	if extra == "" {
		return nil
	}
	m.Channels = make([]rune, len(extra))
	for i := range extra {
		m.Channels[i] = rune(extra[i])
	}
	return nil
}
