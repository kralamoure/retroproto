// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsJoinFriend struct{}

func (m FriendsJoinFriend) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsJoinFriend
}

func (m FriendsJoinFriend) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsJoinFriend) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}