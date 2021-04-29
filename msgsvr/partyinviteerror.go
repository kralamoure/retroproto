// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyInviteError struct{}

func (m PartyInviteError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyInviteError
}

func (m PartyInviteError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyInviteError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
