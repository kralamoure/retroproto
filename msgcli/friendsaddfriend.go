// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsAddFriend struct{}

func NewFriendsAddFriend(extra string) (FriendsAddFriend, error) {
	var m FriendsAddFriend

	if err := m.Deserialize(extra); err != nil {
		return FriendsAddFriend{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsAddFriend) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsAddFriend
}

func (m FriendsAddFriend) MessageName() string {
	return "FriendsAddFriend"
}

func (m FriendsAddFriend) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsAddFriend) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
