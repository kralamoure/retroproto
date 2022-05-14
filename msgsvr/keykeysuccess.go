// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type KeyKeySuccess struct{}

func NewKeyKeySuccess(extra string) (KeyKeySuccess, error) {
	var m KeyKeySuccess

	if err := m.Deserialize(extra); err != nil {
		return KeyKeySuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m KeyKeySuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.KeyKeySuccess
}

func (m KeyKeySuccess) MessageName() string {
	return "KeyKeySuccess"
}

func (m KeyKeySuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyKeySuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
