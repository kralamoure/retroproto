// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsJoinFriend struct{}

func (m FriendsJoinFriend) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsJoinFriend
}

func (m FriendsJoinFriend) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsJoinFriend) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
