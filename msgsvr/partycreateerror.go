// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyCreateError struct{}

func (m PartyCreateError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyCreateError
}

func (m PartyCreateError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyCreateError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
