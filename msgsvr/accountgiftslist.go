// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGiftsList struct{}

func NewAccountGiftsList(extra string) (AccountGiftsList, error) {
	var m AccountGiftsList

	if err := m.Deserialize(extra); err != nil {
		return AccountGiftsList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGiftsList) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountGiftsList
}

func (m AccountGiftsList) MessageName() string {
	return "AccountGiftsList"
}

func (m AccountGiftsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountGiftsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
