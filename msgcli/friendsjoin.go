// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsJoin struct{}

func (m FriendsJoin) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsJoin
}

func (m FriendsJoin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsJoin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
