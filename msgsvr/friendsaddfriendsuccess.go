// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsAddFriendSuccess struct{}

func (m FriendsAddFriendSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsAddFriendSuccess
}

func (m FriendsAddFriendSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsAddFriendSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
