// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeReplayCraft struct{}

func (m ExchangeReplayCraft) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeReplayCraft
}

func (m ExchangeReplayCraft) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeReplayCraft) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
