// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyWhere struct{}

func (m PartyWhere) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyWhere
}

func (m PartyWhere) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyWhere) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
