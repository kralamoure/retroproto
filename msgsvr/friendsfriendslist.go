// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsFriendsList struct{}

func (m FriendsFriendsList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsFriendsList
}

func (m FriendsFriendsList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsFriendsList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
