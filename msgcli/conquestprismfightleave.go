// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightLeave struct{}

func NewConquestPrismFightLeave(extra string) (ConquestPrismFightLeave, error) {
	var m ConquestPrismFightLeave

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismFightLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismFightLeave) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestPrismFightLeave
}

func (m ConquestPrismFightLeave) MessageName() string {
	return "ConquestPrismFightLeave"
}

func (m ConquestPrismFightLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
