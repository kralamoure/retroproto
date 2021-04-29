// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyInviteSuccess struct{}

func (m PartyInviteSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyInviteSuccess
}

func (m PartyInviteSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyInviteSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
