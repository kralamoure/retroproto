// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyAcceptInvitation struct{}

func (m PartyAcceptInvitation) MessageId() retroproto.MsgCliId {
	return retroproto.PartyAcceptInvitation
}

func (m PartyAcceptInvitation) MessageName() string {
	return "PartyAcceptInvitation"
}

func (m PartyAcceptInvitation) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyAcceptInvitation) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
