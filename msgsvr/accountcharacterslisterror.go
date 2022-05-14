// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharactersListError struct{}

func NewAccountCharactersListError(extra string) (AccountCharactersListError, error) {
	var m AccountCharactersListError

	if err := m.Deserialize(extra); err != nil {
		return AccountCharactersListError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharactersListError) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharactersListError
}

func (m AccountCharactersListError) MessageName() string {
	return "AccountCharactersListError"
}

func (m AccountCharactersListError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountCharactersListError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
