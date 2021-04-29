// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type PartyAccept struct{}

func (m PartyAccept) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.PartyAccept
}

func (m PartyAccept) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyAccept) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
