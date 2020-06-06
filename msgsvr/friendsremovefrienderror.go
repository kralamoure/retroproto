// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsRemoveFriendError struct{}

func (m FriendsRemoveFriendError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsRemoveFriendError
}

func (m FriendsRemoveFriendError) Serialized() (string, error) {
	return "", nil
}

func (m *FriendsRemoveFriendError) Deserialize(extra string) error {
	return nil
}
