// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsJoin struct{}

func (m FriendsJoin) ProtocolId() retroproto.MsgCliId {
	return retroproto.FriendsJoin
}

func (m FriendsJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
