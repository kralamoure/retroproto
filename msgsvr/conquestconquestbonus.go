// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestConquestBonus struct{}

func (m ConquestConquestBonus) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestConquestBonus
}

func (m ConquestConquestBonus) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestConquestBonus) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
