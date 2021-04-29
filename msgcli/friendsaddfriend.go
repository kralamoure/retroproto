// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsAddFriend struct{}

func (m FriendsAddFriend) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsAddFriend
}

func (m FriendsAddFriend) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsAddFriend) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
