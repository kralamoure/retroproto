// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyRefuse struct{}

func (m PartyRefuse) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyRefuse
}

func (m PartyRefuse) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyRefuse) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
