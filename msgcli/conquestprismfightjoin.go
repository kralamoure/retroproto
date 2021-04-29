// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismFightJoin struct{}

func (m ConquestPrismFightJoin) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestPrismFightJoin
}

func (m ConquestPrismFightJoin) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismFightJoin) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
