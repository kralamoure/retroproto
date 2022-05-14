// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsRemoveFriendSuccess struct{}

func (m FriendsRemoveFriendSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsRemoveFriendSuccess
}

func (m FriendsRemoveFriendSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsRemoveFriendSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
