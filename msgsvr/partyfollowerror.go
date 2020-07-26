// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyFollowError struct{}

func (m PartyFollowError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyFollowError
}

func (m PartyFollowError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyFollowError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
