// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeKillMount struct{}

func NewExchangeKillMount(extra string) (ExchangeKillMount, error) {
	var m ExchangeKillMount

	if err := m.Deserialize(extra); err != nil {
		return ExchangeKillMount{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeKillMount) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeKillMount
}

func (m ExchangeKillMount) MessageName() string {
	return "ExchangeKillMount"
}

func (m ExchangeKillMount) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeKillMount) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
