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
	return "", nil
}

func (m *PartyLeader) Deserialize(extra string) error {
	return nil
}
