// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EmotesAdd struct{}

func NewEmotesAdd(extra string) (EmotesAdd, error) {
	var m EmotesAdd

	if err := m.Deserialize(extra); err != nil {
		return EmotesAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EmotesAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesAdd
}

func (m EmotesAdd) MessageName() string {
	return "EmotesAdd"
}

func (m EmotesAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
