// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestRequestBalance struct{}

func (m ConquestRequestBalance) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestRequestBalance
}

func (m ConquestRequestBalance) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestRequestBalance) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
