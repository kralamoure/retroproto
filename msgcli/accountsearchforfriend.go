package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSearchForFriend struct {
	Pseudo string
}

func (m AccountSearchForFriend) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountSearchForFriend
}

func (m AccountSearchForFriend) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Pseudo), nil
}

func (m *AccountSearchForFriend) Deserialize(extra string) error {
	m.Pseudo = extra

	return nil
}
