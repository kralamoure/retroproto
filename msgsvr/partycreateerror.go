// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyCreateError struct{}

func (m PartyCreateError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyCreateError
}

func (m PartyCreateError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyCreateError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
