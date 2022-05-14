package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountGetServersList struct{}

func NewAccountGetServersList(extra string) (AccountGetServersList, error) {
	var m AccountGetServersList

	if err := m.Deserialize(extra); err != nil {
		return AccountGetServersList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountGetServersList) MessageId() retroproto.MsgCliId {
	return retroproto.AccountGetServersList
}

func (m AccountGetServersList) MessageName() string {
	return "AccountGetServersList"
}

func (m AccountGetServersList) Serialized() (string, error) {
	return "", nil
}

func (m *AccountGetServersList) Deserialize(extra string) error {
	return nil
}
