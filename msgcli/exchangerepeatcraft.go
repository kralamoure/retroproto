// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeRepeatCraft struct{}

func (m ExchangeRepeatCraft) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeRepeatCraft
}

func (m ExchangeRepeatCraft) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeRepeatCraft) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
