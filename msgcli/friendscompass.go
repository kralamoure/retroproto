// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsCompass struct{}

func (m FriendsCompass) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FriendsCompass
}

func (m FriendsCompass) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsCompass) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
