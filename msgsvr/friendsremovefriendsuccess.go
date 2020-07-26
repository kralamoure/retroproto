// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsRemoveFriendSuccess struct{}

func (m FriendsRemoveFriendSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsRemoveFriendSuccess
}

func (m FriendsRemoveFriendSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsRemoveFriendSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
