// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsFriendsList struct{}

func (m FriendsFriendsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsFriendsList
}

func (m FriendsFriendsList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsFriendsList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
