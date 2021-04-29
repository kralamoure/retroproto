// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsAddFriend struct{}

func (m FriendsAddFriend) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsAddFriend
}

func (m FriendsAddFriend) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsAddFriend) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
