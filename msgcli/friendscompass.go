// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FriendsCompass struct{}

func (m FriendsCompass) ProtocolId() retroproto.MsgCliId {
	return retroproto.FriendsCompass
}

func (m FriendsCompass) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsCompass) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
