// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsAddFriend struct{}

func (m FriendsAddFriend) ProtocolId() retroproto.MsgCliId {
	return retroproto.FriendsAddFriend
}

func (m FriendsAddFriend) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsAddFriend) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
