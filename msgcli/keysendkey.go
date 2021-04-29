// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type KeySendKey struct{}

func (m KeySendKey) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.KeySendKey
}

func (m KeySendKey) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *KeySendKey) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
