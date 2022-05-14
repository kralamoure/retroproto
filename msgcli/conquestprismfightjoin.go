// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightJoin struct{}

func (m ConquestPrismFightJoin) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestPrismFightJoin
}

func (m ConquestPrismFightJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
