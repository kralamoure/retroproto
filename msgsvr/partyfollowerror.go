// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyFollowError struct{}

func (m PartyFollowError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.PartyFollowError
}

func (m PartyFollowError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyFollowError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
