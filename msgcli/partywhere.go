// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyWhere struct{}

func (m PartyWhere) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyWhere
}

func (m PartyWhere) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyWhere) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
