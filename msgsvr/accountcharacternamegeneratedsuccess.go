package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountCharacterNameGeneratedSuccess struct {
	Name string
}

func NewAccountCharacterNameGeneratedSuccess(extra string) (AccountCharacterNameGeneratedSuccess, error) {
	var m AccountCharacterNameGeneratedSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterNameGeneratedSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterNameGeneratedSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterNameGeneratedSuccess
}

func (m AccountCharacterNameGeneratedSuccess) MessageName() string {
	return "AccountCharacterNameGeneratedSuccess"
}

func (m AccountCharacterNameGeneratedSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Name), nil
}

func (m *AccountCharacterNameGeneratedSuccess) Deserialize(extra string) error {
	m.Name = extra

	return nil
}
