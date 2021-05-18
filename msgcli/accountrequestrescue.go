// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountRequestRescue struct{}

func (m AccountRequestRescue) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountRequestRescue
}

func (m AccountRequestRescue) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountRequestRescue) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
