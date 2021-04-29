// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsRemoveFriend struct{}

func (m FriendsRemoveFriend) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsRemoveFriend
}

func (m FriendsRemoveFriend) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsRemoveFriend) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
