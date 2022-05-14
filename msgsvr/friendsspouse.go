// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type FriendsSpouse struct{}

func (m FriendsSpouse) MessageId() retroproto.MsgSvrId {
	return retroproto.FriendsSpouse
}

func (m FriendsSpouse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FriendsSpouse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
