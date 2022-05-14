// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsAddFriendError struct{}

func NewFriendsAddFriendError(extra string) (FriendsAddFriendError, error) {
	var m FriendsAddFriendError

	if err := m.Deserialize(extra); err != nil {
		return FriendsAddFriendError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsAddFriendError) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsAddFriendError
}

func (m FriendsAddFriendError) MessageName() string {
	return "FriendsAddFriendError"
}

func (m FriendsAddFriendError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsAddFriendError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
