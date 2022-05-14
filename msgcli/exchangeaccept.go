// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeAccept struct{}

func NewExchangeAccept(extra string) (ExchangeAccept, error) {
	var m ExchangeAccept

	if err := m.Deserialize(extra); err != nil {
		return ExchangeAccept{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeAccept) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeAccept
}

func (m ExchangeAccept) MessageName() string {
	return "ExchangeAccept"
}

func (m ExchangeAccept) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeAccept) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
