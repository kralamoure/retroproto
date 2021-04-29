// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyFollowSuccess struct{}

func (m PartyFollowSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyFollowSuccess
}

func (m PartyFollowSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyFollowSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
