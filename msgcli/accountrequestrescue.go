// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountRequestRescue struct{}

func NewAccountRequestRescue(extra string) (AccountRequestRescue, error) {
	var m AccountRequestRescue

	if err := m.Deserialize(extra); err != nil {
		return AccountRequestRescue{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountRequestRescue) MessageId() retroproto.MsgCliId {
	return retroproto.AccountRequestRescue
}

func (m AccountRequestRescue) MessageName() string {
	return "AccountRequestRescue"
}

func (m AccountRequestRescue) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountRequestRescue) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
