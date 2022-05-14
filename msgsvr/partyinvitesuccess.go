// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyInviteSuccess struct{}

func (m PartyInviteSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyInviteSuccess
}

func (m PartyInviteSuccess) MessageName() string {
	return "PartyInviteSuccess"
}

func (m PartyInviteSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInviteSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
