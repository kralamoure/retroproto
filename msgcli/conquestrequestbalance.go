// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestRequestBalance struct{}

func (m ConquestRequestBalance) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestRequestBalance
}

func (m ConquestRequestBalance) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestRequestBalance) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
