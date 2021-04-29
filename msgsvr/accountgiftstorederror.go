// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type AccountGiftStoredError struct{}

func (m AccountGiftStoredError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountGiftStoredError
}

func (m AccountGiftStoredError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AccountGiftStoredError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
