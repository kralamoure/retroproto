// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyAccept struct{}

func (m PartyAccept) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyAccept
}

func (m PartyAccept) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyAccept) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
