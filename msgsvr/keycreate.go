// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type KeyCreate struct{}

func NewKeyCreate(extra string) (KeyCreate, error) {
	var m KeyCreate

	if err := m.Deserialize(extra); err != nil {
		return KeyCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m KeyCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.KeyCreate
}

func (m KeyCreate) MessageName() string {
	return "KeyCreate"
}

func (m KeyCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeyCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
