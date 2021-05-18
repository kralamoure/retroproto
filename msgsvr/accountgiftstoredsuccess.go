// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type AccountGiftStoredSuccess struct{}

func (m AccountGiftStoredSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountGiftStoredSuccess
}

func (m AccountGiftStoredSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGiftStoredSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
