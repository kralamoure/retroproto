// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsSetNotifyWhenConnect struct{}

func (m FriendsSetNotifyWhenConnect) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsSetNotifyWhenConnect
}

func (m FriendsSetNotifyWhenConnect) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsSetNotifyWhenConnect) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
