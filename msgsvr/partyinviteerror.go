// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyInviteError struct{}

func (m PartyInviteError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.PartyInviteError
}

func (m PartyInviteError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInviteError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
