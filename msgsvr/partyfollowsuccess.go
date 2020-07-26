// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyFollowSuccess struct{}

func (m PartyFollowSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyFollowSuccess
}

func (m PartyFollowSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyFollowSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
