// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsCompass struct{}

func NewFriendsCompass(extra string) (FriendsCompass, error) {
	var m FriendsCompass

	if err := m.Deserialize(extra); err != nil {
		return FriendsCompass{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsCompass) MessageId() retroproto.MsgCliId {
	return retroproto.FriendsCompass
}

func (m FriendsCompass) MessageName() string {
	return "FriendsCompass"
}

func (m FriendsCompass) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsCompass) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
