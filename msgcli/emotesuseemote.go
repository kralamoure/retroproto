// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type EmotesUseEmote struct{}

func (m EmotesUseEmote) ProtocolId() d1proto.MsgCliId {
	return d1proto.EmotesUseEmote
}

func (m EmotesUseEmote) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *EmotesUseEmote) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
