// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type KeyLeave struct{}

func NewKeyLeave(extra string) (KeyLeave, error) {
	var m KeyLeave

	if err := m.Deserialize(extra); err != nil {
		return KeyLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m KeyLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.KeyLeave
}

func (m KeyLeave) MessageName() string {
	return "KeyLeave"
}

func (m KeyLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
