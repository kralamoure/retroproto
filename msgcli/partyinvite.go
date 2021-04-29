// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyInvite struct{}

func (m PartyInvite) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyInvite
}

func (m PartyInvite) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyInvite) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
