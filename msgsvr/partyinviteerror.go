// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyInviteError struct{}

func (m PartyInviteError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyInviteError
}

func (m PartyInviteError) Serialized() (string, error) {
	return "", nil
}

func (m *PartyInviteError) Deserialize(extra string) error {
	return nil
}
