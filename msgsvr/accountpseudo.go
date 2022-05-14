package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountPseudo struct {
	Value string
}

func NewAccountPseudo(extra string) (AccountPseudo, error) {
	var m AccountPseudo

	if err := m.Deserialize(extra); err != nil {
		return AccountPseudo{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountPseudo) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountPseudo
}

func (m AccountPseudo) MessageName() string {
	return "AccountPseudo"
}

func (m AccountPseudo) Serialized() (string, error) {
	return m.Value, nil
}

func (m *AccountPseudo) Deserialize(extra string) error {
	m.Value = extra

	return nil
}
