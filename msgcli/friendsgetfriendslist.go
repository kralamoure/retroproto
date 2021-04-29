// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsGetFriendsList struct{}

func (m FriendsGetFriendsList) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsGetFriendsList
}

func (m FriendsGetFriendsList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsGetFriendsList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
