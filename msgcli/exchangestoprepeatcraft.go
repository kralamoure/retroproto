// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeStopRepeatCraft struct{}

func NewExchangeStopRepeatCraft(extra string) (ExchangeStopRepeatCraft, error) {
	var m ExchangeStopRepeatCraft

	if err := m.Deserialize(extra); err != nil {
		return ExchangeStopRepeatCraft{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
