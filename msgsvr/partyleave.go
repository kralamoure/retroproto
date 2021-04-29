// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyLeave struct{}

func (m PartyLeave) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyLeave
}

func (m PartyLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
