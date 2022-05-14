// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsSetNotifyWhenConnect struct{}

func (m FriendsSetNotifyWhenConnect) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsSetNotifyWhenConnect
}

func (m FriendsSetNotifyWhenConnect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsSetNotifyWhenConnect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
