// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsGetFriendsList struct{}

func (m FriendsGetFriendsList) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsGetFriendsList
}

func (m FriendsGetFriendsList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsGetFriendsList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
