// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeReplayCraft struct{}

func NewExchangeReplayCraft(extra string) (ExchangeReplayCraft, error) {
	var m ExchangeReplayCraft

	if err := m.Deserialize(extra); err != nil {
		return ExchangeReplayCraft{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeReplayCraft) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeReplayCraft
}

func (m ExchangeReplayCraft) MessageName() string {
	return "ExchangeReplayCraft"
}

func (m ExchangeReplayCraft) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeReplayCraft) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
