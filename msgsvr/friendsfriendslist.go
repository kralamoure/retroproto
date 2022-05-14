// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsFriendsList struct{}

func NewFriendsFriendsList(extra string) (FriendsFriendsList, error) {
	var m FriendsFriendsList

	if err := m.Deserialize(extra); err != nil {
		return FriendsFriendsList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsFriendsList) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsFriendsList
}

func (m FriendsFriendsList) MessageName() string {
	return "FriendsFriendsList"
}

func (m FriendsFriendsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsFriendsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
