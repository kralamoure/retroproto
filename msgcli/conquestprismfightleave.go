// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismFightLeave struct{}

func (m ConquestPrismFightLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestPrismFightLeave
}

func (m ConquestPrismFightLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismFightLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
