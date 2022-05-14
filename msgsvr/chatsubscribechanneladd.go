package msgsvr

import (
	"strings"

	"github.com/kralamoure/retroproto"
)

type ChatSubscribeChannelAdd struct {
	Channels []rune
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
