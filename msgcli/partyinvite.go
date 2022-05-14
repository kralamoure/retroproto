// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyInvite struct{}

func (m PartyInvite) MessageId() retroproto.MsgCliId {
	return retroproto.PartyInvite
}

func (m PartyInvite) MessageName() string {
	return "PartyInvite"
}

func (m PartyInvite) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInvite) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
