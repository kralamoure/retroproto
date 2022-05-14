// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EmotesUseError struct{}

func NewEmotesUseError(extra string) (EmotesUseError, error) {
	var m EmotesUseError

	if err := m.Deserialize(extra); err != nil {
		return EmotesUseError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EmotesUseError) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesUseError
}

func (m EmotesUseError) MessageName() string {
	return "EmotesUseError"
}

func (m EmotesUseError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesUseError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
