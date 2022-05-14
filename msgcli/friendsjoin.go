// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsJoin struct{}

func NewFriendsJoin(extra string) (FriendsJoin, error) {
	var m FriendsJoin

	if err := m.Deserialize(extra); err != nil {
		return FriendsJoin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsJoin) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsJoin
}

func (m FriendsJoin) MessageName() string {
	return "FriendsJoin"
}

func (m FriendsJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
