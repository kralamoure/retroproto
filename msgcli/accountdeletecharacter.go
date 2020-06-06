// TODO
package msgcli

import (
	"fmt"
	"html"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type AccountDeleteCharacter struct {
	Id           int
	SecretAnswer string
}

func (m AccountDeleteCharacter) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountDeleteCharacter
}

func (m AccountDeleteCharacter) Serialized() (string, error) {
	return fmt.Sprintf("%d|%s", m.Id, m.SecretAnswer), nil
}

func (m *AccountDeleteCharacter) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return d1proto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	m.SecretAnswer = html.UnescapeString(sli[1])

	return nil
}
