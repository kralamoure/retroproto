// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsJoin struct{}

func (m FriendsJoin) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsJoin
}

func (m FriendsJoin) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsJoin) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
