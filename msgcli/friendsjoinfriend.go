// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsJoinFriend struct{}

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
