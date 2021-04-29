package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsNotifyChange struct {
	Notify bool
}

func (m FriendsNotifyChange) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsNotifyChange
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
