// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type KeySendKey struct{}

func NewKeySendKey(extra string) (KeySendKey, error) {
	var m KeySendKey

	if err := m.Deserialize(extra); err != nil {
		return KeySendKey{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m KeySendKey) MessageId() retroproto.MsgCliId {
	return retroproto.KeySendKey
}

func (m KeySendKey) MessageName() string {
	return "KeySendKey"
}

func (m KeySendKey) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *KeySendKey) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
