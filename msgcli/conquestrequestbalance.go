// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestRequestBalance struct{}

func NewConquestRequestBalance(extra string) (ConquestRequestBalance, error) {
	var m ConquestRequestBalance

	if err := m.Deserialize(extra); err != nil {
		return ConquestRequestBalance{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestRequestBalance) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestRequestBalance
}

func (m ConquestRequestBalance) MessageName() string {
	return "ConquestRequestBalance"
}

func (m ConquestRequestBalance) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestRequestBalance) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
