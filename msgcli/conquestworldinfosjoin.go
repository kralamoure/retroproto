// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestWorldInfosJoin struct{}

func (m ConquestWorldInfosJoin) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestWorldInfosJoin
}

func (m ConquestWorldInfosJoin) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestWorldInfosJoin) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
