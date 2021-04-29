// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGiftStoredSuccess struct{}

func (m AccountGiftStoredSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountGiftStoredSuccess
}

func (m AccountGiftStoredSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountGiftStoredSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
