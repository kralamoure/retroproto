// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestGetAlignedBonus struct{}

func (m ConquestGetAlignedBonus) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestGetAlignedBonus
}

func (m ConquestGetAlignedBonus) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestGetAlignedBonus) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
