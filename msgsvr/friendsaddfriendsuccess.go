// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsAddFriendSuccess struct{}

func NewFriendsAddFriendSuccess(extra string) (FriendsAddFriendSuccess, error) {
	var m FriendsAddFriendSuccess

	if err := m.Deserialize(extra); err != nil {
		return FriendsAddFriendSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsAddFriendSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsAddFriendSuccess
}

func (m FriendsAddFriendSuccess) MessageName() string {
	return "FriendsAddFriendSuccess"
}

func (m FriendsAddFriendSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsAddFriendSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
