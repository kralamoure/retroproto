package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsNotifyChange struct {
	Notify bool
}

func (m FriendsNotifyChange) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsNotifyChange
}

func (m FriendsNotifyChange) MessageName() string {
	return "FriendsNotifyChange"
}

func (m FriendsNotifyChange) Serialized() (string, error) {
	notify := "-"
	if m.Notify {
		notify = "+"
	}

	return notify, nil
}

func (m *FriendsNotifyChange) Deserialize(extra string) error {
	m.Notify = extra == "+"

	return nil
}
