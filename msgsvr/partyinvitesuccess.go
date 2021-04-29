// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyInviteSuccess struct{}

func (m PartyInviteSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyInviteSuccess
}

func (m PartyInviteSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyInviteSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
