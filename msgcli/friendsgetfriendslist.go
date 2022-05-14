// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsGetFriendsList struct{}

func NewFriendsGetFriendsList(extra string) (FriendsGetFriendsList, error) {
	var m FriendsGetFriendsList

	if err := m.Deserialize(extra); err != nil {
		return FriendsGetFriendsList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsGetFriendsList) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsGetFriendsList
}

func (m FriendsGetFriendsList) MessageName() string {
	return "FriendsGetFriendsList"
}

func (m FriendsGetFriendsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsGetFriendsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
