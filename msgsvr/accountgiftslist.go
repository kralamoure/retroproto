// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountGiftsList struct{}

func (m AccountGiftsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountGiftsList
}

func (m AccountGiftsList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *AccountGiftsList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
