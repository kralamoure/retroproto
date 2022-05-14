// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddEnemyRemove struct{}

func NewConquestPrismFightAddEnemyRemove(extra string) (ConquestPrismFightAddEnemyRemove, error) {
	var m ConquestPrismFightAddEnemyRemove

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismFightAddEnemyRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismFightAddEnemyRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddEnemyRemove
}

func (m ConquestPrismFightAddEnemyRemove) MessageName() string {
	return "ConquestPrismFightAddEnemyRemove"
}

func (m ConquestPrismFightAddEnemyRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddEnemyRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
