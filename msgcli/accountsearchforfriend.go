package msgcli

import (
	"fmt"

	"github.com/kralamoure/d1encoding"
)

type AccountSearchForFriend struct {
	Pseudo string
}

func (m AccountSearchForFriend) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountSearchForFriend
}

func (m AccountSearchForFriend) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Pseudo), nil
}

func (m *AccountSearchForFriend) Deserialize(extra string) error {
	m.Pseudo = extra

	return nil
}
