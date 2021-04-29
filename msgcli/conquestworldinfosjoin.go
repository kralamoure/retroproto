// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestWorldInfosJoin struct{}

func (m ConquestWorldInfosJoin) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestWorldInfosJoin
}

func (m ConquestWorldInfosJoin) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestWorldInfosJoin) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
