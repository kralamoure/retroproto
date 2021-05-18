// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountGiftStoredError struct{}

func (m AccountGiftStoredError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountGiftStoredError
}

func (m AccountGiftStoredError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGiftStoredError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
