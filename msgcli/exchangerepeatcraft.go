// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeRepeatCraft struct{}

func NewExchangeRepeatCraft(extra string) (ExchangeRepeatCraft, error) {
	var m ExchangeRepeatCraft

	if err := m.Deserialize(extra); err != nil {
		return ExchangeRepeatCraft{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeRepeatCraft) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeRepeatCraft
}

func (m ExchangeRepeatCraft) MessageName() string {
	return "ExchangeRepeatCraft"
}

func (m ExchangeRepeatCraft) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRepeatCraft) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
