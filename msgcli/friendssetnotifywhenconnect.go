// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsSetNotifyWhenConnect struct{}

func (m FriendsSetNotifyWhenConnect) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsSetNotifyWhenConnect
}

func (m FriendsSetNotifyWhenConnect) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsSetNotifyWhenConnect) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
