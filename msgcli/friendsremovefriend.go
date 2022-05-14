// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsRemoveFriend struct{}

func NewFriendsRemoveFriend(extra string) (FriendsRemoveFriend, error) {
	var m FriendsRemoveFriend

	if err := m.Deserialize(extra); err != nil {
		return FriendsRemoveFriend{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsRemoveFriend) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsRemoveFriend
}

func (m FriendsRemoveFriend) MessageName() string {
	return "FriendsRemoveFriend"
}

func (m FriendsRemoveFriend) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsRemoveFriend) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
