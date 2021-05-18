// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyRefuse struct{}

func (m PartyRefuse) ProtocolId() retroproto.MsgSvrId {
	return retroproto.PartyRefuse
}

func (m PartyRefuse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRefuse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
