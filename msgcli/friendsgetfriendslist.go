// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsGetFriendsList struct{}

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
