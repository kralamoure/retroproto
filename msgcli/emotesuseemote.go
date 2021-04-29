// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type EmotesUseEmote struct{}

func (m EmotesUseEmote) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.EmotesUseEmote
}

func (m EmotesUseEmote) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EmotesUseEmote) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
