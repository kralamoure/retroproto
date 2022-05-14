// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyLeader struct{}

func (m PartyLeader) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyLeader
}

func (m PartyLeader) MessageName() string {
	return "PartyLeader"
}

func (m PartyLeader) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyLeader) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
