// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsFriendsList struct{}

func (m FriendsFriendsList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.FriendsFriendsList
}

func (m FriendsFriendsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsFriendsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
