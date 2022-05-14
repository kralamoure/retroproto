// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeKillMountInPark struct{}

func NewExchangeKillMountInPark(extra string) (ExchangeKillMountInPark, error) {
	var m ExchangeKillMountInPark

	if err := m.Deserialize(extra); err != nil {
		return ExchangeKillMountInPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeKillMountInPark) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeKillMountInPark
}

func (m ExchangeKillMountInPark) MessageName() string {
	return "ExchangeKillMountInPark"
}

func (m ExchangeKillMountInPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeKillMountInPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
