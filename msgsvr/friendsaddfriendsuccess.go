// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsAddFriendSuccess struct{}

func (m FriendsAddFriendSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.FriendsAddFriendSuccess
}

func (m FriendsAddFriendSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsAddFriendSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
