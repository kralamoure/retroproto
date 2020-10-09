// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsRemoveFriend struct{}

func (m FriendsRemoveFriend) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsRemoveFriend
}

func (m FriendsRemoveFriend) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsRemoveFriend) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
