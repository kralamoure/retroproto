package msgsvr

import (
	"strings"

	"github.com/kralamoure/d1proto"
)

type ChatSubscribeChannelAdd struct {
	Channels []rune
}

func (m ChatSubscribeChannelAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ChatSubscribeChannelAdd
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
