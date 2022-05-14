// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsAddFriend struct{}

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
