// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ExchangeStopRepeatCraft struct{}

func (m ExchangeStopRepeatCraft) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeStopRepeatCraft
}

func (m ExchangeStopRepeatCraft) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ExchangeStopRepeatCraft) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
