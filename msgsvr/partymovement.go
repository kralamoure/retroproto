// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyMovement struct{}

func (m PartyMovement) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyMovement
}

func (m PartyMovement) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyMovement) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
