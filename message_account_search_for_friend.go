package retroproto

import (
	"fmt"

	"github.com/kralamoure/retropb/gen/go/retropb"
)

type AccountSearchForFriend struct {
	*retropb.AccountSearchForFriend
}

func NewAccountSearchForFriend(extra string) (AccountSearchForFriend, error) {
	var m retropb.AccountSearchForFriend

	m.Nickname = extra

	return AccountSearchForFriend{AccountSearchForFriend: &m}, nil
}

func (m AccountSearchForFriend) MessageId() string {
	return "AF"
}

func (m AccountSearchForFriend) String() string {
	return fmt.Sprintf("%s", m.GetNickname())
}
