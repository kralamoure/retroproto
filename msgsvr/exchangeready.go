// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeReady struct{}

func NewExchangeReady(extra string) (ExchangeReady, error) {
	var m ExchangeReady

	if err := m.Deserialize(extra); err != nil {
		return ExchangeReady{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeReady) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeReady
}

func (m ExchangeReady) MessageName() string {
	return "ExchangeReady"
}

func (m ExchangeReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
