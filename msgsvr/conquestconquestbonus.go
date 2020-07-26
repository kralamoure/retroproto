// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestConquestBonus struct{}

func (m ConquestConquestBonus) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestConquestBonus
}

func (m ConquestConquestBonus) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestConquestBonus) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
