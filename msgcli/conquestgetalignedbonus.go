// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestGetAlignedBonus struct{}

func (m ConquestGetAlignedBonus) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestGetAlignedBonus
}

func (m ConquestGetAlignedBonus) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestGetAlignedBonus) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
