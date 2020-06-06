// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountRequestRescue struct{}

func (m AccountRequestRescue) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountRequestRescue
}

func (m AccountRequestRescue) Serialized() (string, error) {
	return "", nil
}

func (m *AccountRequestRescue) Deserialize(extra string) error {
	return nil
}
