// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyLeader struct{}

func (m PartyLeader) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyLeader
}

func (m PartyLeader) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyLeader) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
