// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FriendsSpouse struct{}

func NewFriendsSpouse(extra string) (FriendsSpouse, error) {
	var m FriendsSpouse

	if err := m.Deserialize(extra); err != nil {
		return FriendsSpouse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FriendsSpouse) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsSpouse
}

func (m FriendsSpouse) MessageName() string {
	return "FriendsSpouse"
}

func (m FriendsSpouse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsSpouse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
