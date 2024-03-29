package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountSetCharacter struct {
	Id int
}

func NewAccountSetCharacter(extra string) (AccountSetCharacter, error) {
	var m AccountSetCharacter

	if err := m.Deserialize(extra); err != nil {
		return AccountSetCharacter{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSetCharacter) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSetCharacter
}

func (m AccountSetCharacter) MessageName() string {
	return "AccountSetCharacter"
}

func (m AccountSetCharacter) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *AccountSetCharacter) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
