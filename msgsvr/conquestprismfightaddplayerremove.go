// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddPlayerRemove struct{}

func NewConquestPrismFightAddPlayerRemove(extra string) (ConquestPrismFightAddPlayerRemove, error) {
	var m ConquestPrismFightAddPlayerRemove

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismFightAddPlayerRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismFightAddPlayerRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddPlayerRemove
}

func (m ConquestPrismFightAddPlayerRemove) MessageName() string {
	return "ConquestPrismFightAddPlayerRemove"
}

func (m ConquestPrismFightAddPlayerRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
