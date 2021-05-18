// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeStopRepeatCraft struct{}

func (m ExchangeStopRepeatCraft) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeStopRepeatCraft
}

func (m ExchangeStopRepeatCraft) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeStopRepeatCraft) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
