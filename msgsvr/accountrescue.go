// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountRescue struct{}

func NewAccountRescue(extra string) (AccountRescue, error) {
	var m AccountRescue

	if err := m.Deserialize(extra); err != nil {
		return AccountRescue{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountRescue) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountRescue
}

func (m AccountRescue) MessageName() string {
	return "AccountRescue"
}

func (m AccountRescue) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountRescue) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
