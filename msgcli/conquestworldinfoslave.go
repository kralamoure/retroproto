// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestWorldInfosLave struct{}

func (m ConquestWorldInfosLave) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestWorldInfosLave
}

func (m ConquestWorldInfosLave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestWorldInfosLave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
