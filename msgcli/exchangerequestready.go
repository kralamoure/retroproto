// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeRequestReady struct{}

func NewExchangeRequestReady(extra string) (ExchangeRequestReady, error) {
	var m ExchangeRequestReady

	if err := m.Deserialize(extra); err != nil {
		return ExchangeRequestReady{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeRequestReady) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeRequestReady
}

func (m ExchangeRequestReady) MessageName() string {
	return "ExchangeRequestReady"
}

func (m ExchangeRequestReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRequestReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
