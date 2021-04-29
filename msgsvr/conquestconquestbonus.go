// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestConquestBonus struct{}

func (m ConquestConquestBonus) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestConquestBonus
}

func (m ConquestConquestBonus) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestConquestBonus) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
