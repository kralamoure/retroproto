package msgsvr

import (
	"strings"

	"github.com/kralamoure/retroproto"
)

type ChatSubscribeChannelRemove struct {
	Channels []rune
}

func (m ChatSubscribeChannelRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ChatSubscribeChannelRemove
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
