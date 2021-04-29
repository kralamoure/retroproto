// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsRemoveFriendError struct{}

func (m FriendsRemoveFriendError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsRemoveFriendError
}

func (m FriendsRemoveFriendError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsRemoveFriendError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
