// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountGiftStoredError struct{}

func (m AccountGiftStoredError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountGiftStoredError
}

func (m AccountGiftStoredError) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGiftStoredError) Deserialize(extra string) error {
	return nil
}
