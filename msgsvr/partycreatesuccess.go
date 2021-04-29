// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyCreateSuccess struct{}

func (m PartyCreateSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyCreateSuccess
}

func (m PartyCreateSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyCreateSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
