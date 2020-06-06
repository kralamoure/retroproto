// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type AccountGiftStoredSuccess struct{}

func (m AccountGiftStoredSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountGiftStoredSuccess
}

func (m AccountGiftStoredSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGiftStoredSuccess) Deserialize(extra string) error {
	return nil
}
