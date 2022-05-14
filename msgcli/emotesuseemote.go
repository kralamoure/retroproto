// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type EmotesUseEmote struct{}

func (m EmotesUseEmote) MessageId() retroproto.MsgCliId {
	return retroproto.EmotesUseEmote
}

func (m EmotesUseEmote) MessageName() string {
	return "EmotesUseEmote"
}

func (m EmotesUseEmote) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesUseEmote) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
