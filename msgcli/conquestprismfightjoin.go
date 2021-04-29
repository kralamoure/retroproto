// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismFightJoin struct{}

func (m ConquestPrismFightJoin) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestPrismFightJoin
}

func (m ConquestPrismFightJoin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismFightJoin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
