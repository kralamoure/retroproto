// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyInvite struct{}

func (m PartyInvite) ProtocolId() retroproto.MsgCliId {
	return retroproto.PartyInvite
}

func (m PartyInvite) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInvite) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
