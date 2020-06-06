// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsAddFriendError struct{}

func (m FriendsAddFriendError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsAddFriendError
}

func (m FriendsAddFriendError) Serialized() (string, error) {
	return "", nil
}

func (m *FriendsAddFriendError) Deserialize(extra string) error {
	return nil
}
