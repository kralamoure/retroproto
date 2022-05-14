// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddEnemyAdd struct{}

func NewConquestPrismFightAddEnemyAdd(extra string) (ConquestPrismFightAddEnemyAdd, error) {
	var m ConquestPrismFightAddEnemyAdd

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismFightAddEnemyAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismFightAddEnemyAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddEnemyAdd
}

func (m ConquestPrismFightAddEnemyAdd) MessageName() string {
	return "ConquestPrismFightAddEnemyAdd"
}

func (m ConquestPrismFightAddEnemyAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddEnemyAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
