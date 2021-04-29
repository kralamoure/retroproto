// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyRefuse struct{}

func (m PartyRefuse) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyRefuse
}

func (m PartyRefuse) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyRefuse) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
