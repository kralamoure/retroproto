// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsRemoveFriendSuccess struct{}

func (m FriendsRemoveFriendSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsRemoveFriendSuccess
}

func (m FriendsRemoveFriendSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsRemoveFriendSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
