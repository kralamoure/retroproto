// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsJoinFriend struct{}

func NewFriendsJoinFriend(extra string) (FriendsJoinFriend, error) {
	var m FriendsJoinFriend

	if err := m.Deserialize(extra); err != nil {
		return FriendsJoinFriend{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsJoinFriend) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsJoinFriend
}

func (m FriendsJoinFriend) MessageName() string {
	return "FriendsJoinFriend"
}

func (m FriendsJoinFriend) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsJoinFriend) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
