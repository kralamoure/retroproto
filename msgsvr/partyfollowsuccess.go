// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyFollowSuccess struct{}

func (m PartyFollowSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyFollowSuccess
}

func (m PartyFollowSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyFollowSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
