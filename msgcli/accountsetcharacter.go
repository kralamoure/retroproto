package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type AccountSetCharacter struct {
	Id int
}

func (m AccountSetCharacter) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountSetCharacter
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
