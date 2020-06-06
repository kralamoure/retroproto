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
	return "", nil
}

func (m *AccountGiftsList) Deserialize(extra string) error {
	return nil
}
