// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLeave struct{}

func NewExchangeLeave(extra string) (ExchangeLeave, error) {
	var m ExchangeLeave

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLeave) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeLeave
}

func (m ExchangeLeave) MessageName() string {
	return "ExchangeLeave"
}

func (m ExchangeLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
