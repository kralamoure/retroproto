// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismInfosJoin struct{}

func (m ConquestPrismInfosJoin) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestPrismInfosJoin
}

func (m ConquestPrismInfosJoin) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismInfosJoin) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}