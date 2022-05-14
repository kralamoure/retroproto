// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyInviteError struct{}

func (m PartyInviteError) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyInviteError
}

func (m PartyInviteError) MessageName() string {
	return "PartyInviteError"
}

func (m PartyInviteError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInviteError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
