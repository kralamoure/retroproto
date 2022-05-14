// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsSetNotifyWhenConnect struct{}

func NewFriendsSetNotifyWhenConnect(extra string) (FriendsSetNotifyWhenConnect, error) {
	var m FriendsSetNotifyWhenConnect

	if err := m.Deserialize(extra); err != nil {
		return FriendsSetNotifyWhenConnect{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsSetNotifyWhenConnect) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsSetNotifyWhenConnect
}

func (m FriendsSetNotifyWhenConnect) MessageName() string {
	return "FriendsSetNotifyWhenConnect"
}

func (m FriendsSetNotifyWhenConnect) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsSetNotifyWhenConnect) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
