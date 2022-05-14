// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosLeave struct{}

func (m ConquestPrismInfosLeave) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestPrismInfosLeave
}

func (m ConquestPrismInfosLeave) MessageName() string {
	return "ConquestPrismInfosLeave"
}

func (m ConquestPrismInfosLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
