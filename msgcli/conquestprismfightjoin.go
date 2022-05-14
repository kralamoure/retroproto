// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightJoin struct{}

func NewConquestPrismFightJoin(extra string) (ConquestPrismFightJoin, error) {
	var m ConquestPrismFightJoin

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismFightJoin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismFightJoin) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestPrismFightJoin
}

func (m ConquestPrismFightJoin) MessageName() string {
	return "ConquestPrismFightJoin"
}

func (m ConquestPrismFightJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
