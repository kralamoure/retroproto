package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSendIdentity struct {
	Id string
}

func NewAccountSendIdentity(extra string) (AccountSendIdentity, error) {
	var m AccountSendIdentity

	if err := m.Deserialize(extra); err != nil {
		return AccountSendIdentity{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSendIdentity) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSendIdentity
}

func (m AccountSendIdentity) MessageName() string {
	return "AccountSendIdentity"
}

func (m AccountSendIdentity) Serialized() (string, error) {
	return m.Id, nil
}

func (m *AccountSendIdentity) Deserialize(extra string) error {
	m.Id = extra

	return nil
}
