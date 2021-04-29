// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountRequestRescue struct{}

func (m AccountRequestRescue) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountRequestRescue
}

func (m AccountRequestRescue) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountRequestRescue) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
