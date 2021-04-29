// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyMovement struct{}

func (m PartyMovement) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyMovement
}

func (m PartyMovement) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyMovement) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
