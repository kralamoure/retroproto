// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyLeader struct{}

func (m PartyLeader) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyLeader
}

func (m PartyLeader) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyLeader) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
