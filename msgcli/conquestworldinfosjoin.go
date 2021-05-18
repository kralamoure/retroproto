// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestWorldInfosJoin struct{}

func (m ConquestWorldInfosJoin) ProtocolId() retroproto.MsgCliId {
	return retroproto.ConquestWorldInfosJoin
}

func (m ConquestWorldInfosJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestWorldInfosJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
