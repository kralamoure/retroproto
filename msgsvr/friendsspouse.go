// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type FriendsSpouse struct{}

func (m FriendsSpouse) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.FriendsSpouse
}

func (m FriendsSpouse) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FriendsSpouse) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
