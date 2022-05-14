// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EmotesUseEmote struct{}

func NewEmotesUseEmote(extra string) (EmotesUseEmote, error) {
	var m EmotesUseEmote

	if err := m.Deserialize(extra); err != nil {
		return EmotesUseEmote{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
