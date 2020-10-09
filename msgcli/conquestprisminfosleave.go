// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismInfosLeave struct{}

func (m ConquestPrismInfosLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestPrismInfosLeave
}

func (m ConquestPrismInfosLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismInfosLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
