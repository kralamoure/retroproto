// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyRefuseInvitation struct{}

func (m PartyRefuseInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.PartyRefuseInvitation
}

func (m PartyRefuseInvitation) MessageName() string {
	return "PartyRefuseInvitation"
}

func (m PartyRefuseInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRefuseInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
