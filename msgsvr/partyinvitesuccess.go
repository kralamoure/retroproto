// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyInviteSuccess struct{}

func (m PartyInviteSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.PartyInviteSuccess
}

func (m PartyInviteSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyInviteSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
