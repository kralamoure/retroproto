// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosJoin struct{}

func (m ConquestPrismInfosJoin) ProtocolId() retroproto.MsgCliId {
	return retroproto.ConquestPrismInfosJoin
}

func (m ConquestPrismInfosJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
