// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesRemoveEnemy struct{}

func NewEnemiesRemoveEnemy(extra string) (EnemiesRemoveEnemy, error) {
	var m EnemiesRemoveEnemy

	if err := m.Deserialize(extra); err != nil {
		return EnemiesRemoveEnemy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesRemoveEnemy) MessageId() retroproto.MsgCliId {
	return retroproto.EnemiesRemoveEnemy
}

func (m EnemiesRemoveEnemy) MessageName() string {
	return "EnemiesRemoveEnemy"
}

func (m EnemiesRemoveEnemy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesRemoveEnemy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
