// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyWhere struct{}

func (m PartyWhere) ProtocolId() retroproto.MsgCliId {
	return retroproto.PartyWhere
}

func (m PartyWhere) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyWhere) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
