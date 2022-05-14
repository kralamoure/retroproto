// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EmotesUseSuccess struct{}

func NewEmotesUseSuccess(extra string) (EmotesUseSuccess, error) {
	var m EmotesUseSuccess

	if err := m.Deserialize(extra); err != nil {
		return EmotesUseSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EmotesUseSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesUseSuccess
}

func (m EmotesUseSuccess) MessageName() string {
	return "EmotesUseSuccess"
}

func (m EmotesUseSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesUseSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
