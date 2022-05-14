// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSetNickname struct{}

func NewAccountSetNickname(extra string) (AccountSetNickname, error) {
	var m AccountSetNickname

	if err := m.Deserialize(extra); err != nil {
		return AccountSetNickname{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSetNickname) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSetNickname
}

func (m AccountSetNickname) MessageName() string {
	return "AccountSetNickname"
}

func (m AccountSetNickname) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AccountSetNickname) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
