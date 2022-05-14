// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsRemoveFriendSuccess struct{}

func NewFriendsRemoveFriendSuccess(extra string) (FriendsRemoveFriendSuccess, error) {
	var m FriendsRemoveFriendSuccess

	if err := m.Deserialize(extra); err != nil {
		return FriendsRemoveFriendSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsRemoveFriendSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsRemoveFriendSuccess
}

func (m FriendsRemoveFriendSuccess) MessageName() string {
	return "FriendsRemoveFriendSuccess"
}

func (m FriendsRemoveFriendSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsRemoveFriendSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
