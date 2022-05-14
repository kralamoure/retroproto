// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeSetPublicMode struct{}

func NewExchangeSetPublicMode(extra string) (ExchangeSetPublicMode, error) {
	var m ExchangeSetPublicMode

	if err := m.Deserialize(extra); err != nil {
		return ExchangeSetPublicMode{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeSetPublicMode) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeSetPublicMode
}

func (m ExchangeSetPublicMode) MessageName() string {
	return "ExchangeSetPublicMode"
}

func (m ExchangeSetPublicMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeSetPublicMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
