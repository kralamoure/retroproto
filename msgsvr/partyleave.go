// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyLeave struct{}

func (m PartyLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyLeave
}

func (m PartyLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
