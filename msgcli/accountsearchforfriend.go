package msgcli

import (
	"fmt"

	"github.com/kralamoure/d1proto"
)

type AccountSearchForFriend struct {
	Pseudo string
}

func (m AccountSearchForFriend) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountSearchForFriend
}

func (m AccountSearchForFriend) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Pseudo), nil
}

func (m *AccountSearchForFriend) Deserialize(extra string) error {
	m.Pseudo = extra

	return nil
}
