// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsAddFriendSuccess struct{}

func (m FriendsAddFriendSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsAddFriendSuccess
}

func (m FriendsAddFriendSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *FriendsAddFriendSuccess) Deserialize(extra string) error {
	return nil
}
