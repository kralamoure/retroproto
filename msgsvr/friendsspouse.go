// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type FriendsSpouse struct{}

func (m FriendsSpouse) ProtocolId() d1proto.MsgSvrId {
	return d1proto.FriendsSpouse
}

func (m FriendsSpouse) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FriendsSpouse) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
