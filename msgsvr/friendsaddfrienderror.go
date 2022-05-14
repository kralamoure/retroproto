// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsAddFriendError struct{}

func (m FriendsAddFriendError) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsAddFriendError
}

func (m FriendsAddFriendError) MessageName() string {
	return "FriendsAddFriendError"
}

func (m FriendsAddFriendError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsAddFriendError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
