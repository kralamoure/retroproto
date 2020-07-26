// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismFightLeave struct{}

func (m ConquestPrismFightLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestPrismFightLeave
}

func (m ConquestPrismFightLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismFightLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
