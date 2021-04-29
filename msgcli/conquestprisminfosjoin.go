// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismInfosJoin struct{}

func (m ConquestPrismInfosJoin) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestPrismInfosJoin
}

func (m ConquestPrismInfosJoin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismInfosJoin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
