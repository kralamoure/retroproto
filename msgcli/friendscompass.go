// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FriendsCompass struct{}

func (m FriendsCompass) ProtocolId() d1proto.MsgCliId {
	return d1proto.FriendsCompass
}

func (m FriendsCompass) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsCompass) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
