// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsRemoveFriend struct{}

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
