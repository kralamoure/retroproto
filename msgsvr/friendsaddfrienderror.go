// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsAddFriendError struct{}

func (m FriendsAddFriendError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsAddFriendError
}

func (m FriendsAddFriendError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsAddFriendError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
