// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeStopRepeatCraft struct{}

func (m ExchangeStopRepeatCraft) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeStopRepeatCraft
}

func (m ExchangeStopRepeatCraft) MessageName() string {
	return "ExchangeStopRepeatCraft"
}

func (m ExchangeStopRepeatCraft) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeStopRepeatCraft) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
