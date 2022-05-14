// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EmotesRemove struct{}

func NewEmotesRemove(extra string) (EmotesRemove, error) {
	var m EmotesRemove

	if err := m.Deserialize(extra); err != nil {
		return EmotesRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EmotesRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesRemove
}

func (m EmotesRemove) MessageName() string {
	return "EmotesRemove"
}

func (m EmotesRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
