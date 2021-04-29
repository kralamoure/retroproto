// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type KeySendKey struct{}

func (m KeySendKey) ProtocolId() d1proto.MsgCliId {
	return d1proto.KeySendKey
}

func (m KeySendKey) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *KeySendKey) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
