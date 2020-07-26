// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyCreateSuccess struct{}

func (m PartyCreateSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyCreateSuccess
}

func (m PartyCreateSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyCreateSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
