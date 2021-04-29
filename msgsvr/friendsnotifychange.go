package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsNotifyChange struct {
	Notify bool
}

func (m FriendsNotifyChange) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsNotifyChange
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
