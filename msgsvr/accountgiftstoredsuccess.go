// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGiftStoredSuccess struct{}

func NewAccountGiftStoredSuccess(extra string) (AccountGiftStoredSuccess, error) {
	var m AccountGiftStoredSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountGiftStoredSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGiftStoredSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountGiftStoredSuccess
}

func (m AccountGiftStoredSuccess) MessageName() string {
	return "AccountGiftStoredSuccess"
}

func (m AccountGiftStoredSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGiftStoredSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
