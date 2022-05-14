package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountSearchForFriend struct {
	Pseudo string
}

func NewAccountSearchForFriend(extra string) (AccountSearchForFriend, error) {
	var m AccountSearchForFriend

	if err := m.Deserialize(extra); err != nil {
		return AccountSearchForFriend{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountSearchForFriend) MessageId() retroproto.MsgCliId {
	return retroproto.AccountSearchForFriend
}

func (m AccountSearchForFriend) MessageName() string {
	return "AccountSearchForFriend"
}

func (m AccountSearchForFriend) Serialized() (string, error) {
	return fmt.Sprintf("%s", m.Pseudo), nil
}

func (m *AccountSearchForFriend) Deserialize(extra string) error {
	m.Pseudo = extra

	return nil
}
