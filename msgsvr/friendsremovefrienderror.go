// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsRemoveFriendError struct{}

func NewFriendsRemoveFriendError(extra string) (FriendsRemoveFriendError, error) {
	var m FriendsRemoveFriendError

	if err := m.Deserialize(extra); err != nil {
		return FriendsRemoveFriendError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsRemoveFriendError) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsRemoveFriendError
}

func (m FriendsRemoveFriendError) MessageName() string {
	return "FriendsRemoveFriendError"
}

func (m FriendsRemoveFriendError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsRemoveFriendError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
