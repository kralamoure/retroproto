// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsRemoveFriendError struct{}

func (m FriendsRemoveFriendError) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsRemoveFriendError
}

func (m FriendsRemoveFriendError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsRemoveFriendError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
