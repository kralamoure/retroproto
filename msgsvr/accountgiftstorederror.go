// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGiftStoredError struct{}

func NewAccountGiftStoredError(extra string) (AccountGiftStoredError, error) {
	var m AccountGiftStoredError

	if err := m.Deserialize(extra); err != nil {
		return AccountGiftStoredError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGiftStoredError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountGiftStoredError
}

func (m AccountGiftStoredError) MessageName() string {
	return "AccountGiftStoredError"
}

func (m AccountGiftStoredError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGiftStoredError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
