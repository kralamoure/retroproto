// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyInvite struct{}

func (m PartyInvite) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyInvite
}

func (m PartyInvite) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyInvite) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
