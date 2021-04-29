// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGiftsList struct{}

func (m AccountGiftsList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountGiftsList
}

func (m AccountGiftsList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountGiftsList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
