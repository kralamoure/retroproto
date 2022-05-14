// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountServersListError struct{}

func NewAccountServersListError(extra string) (AccountServersListError, error) {
	var m AccountServersListError

	if err := m.Deserialize(extra); err != nil {
		return AccountServersListError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountServersListError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountServersListError
}

func (m AccountServersListError) MessageName() string {
	return "AccountServersListError"
}

func (m AccountServersListError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountServersListError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
