// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type EmotesUseEmote struct{}

func (m EmotesUseEmote) ProtocolId() retroproto.MsgCliId {
	return retroproto.EmotesUseEmote
}

func (m EmotesUseEmote) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesUseEmote) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
