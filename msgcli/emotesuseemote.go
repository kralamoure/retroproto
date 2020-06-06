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
	return "", nil
}

func (m *EmotesUseEmote) Deserialize(extra string) error {
	return nil
}
