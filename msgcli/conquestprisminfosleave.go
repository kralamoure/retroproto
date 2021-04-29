// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismInfosLeave struct{}

func (m ConquestPrismInfosLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestPrismInfosLeave
}

func (m ConquestPrismInfosLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismInfosLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
