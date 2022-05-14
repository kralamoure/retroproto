// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesAddEnemy struct{}

func NewEnemiesAddEnemy(extra string) (EnemiesAddEnemy, error) {
	var m EnemiesAddEnemy

	if err := m.Deserialize(extra); err != nil {
		return EnemiesAddEnemy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesAddEnemy) MessageId() retroproto.MsgCliId {
	return retroproto.EnemiesAddEnemy
}

func (m EnemiesAddEnemy) MessageName() string {
	return "EnemiesAddEnemy"
}

func (m EnemiesAddEnemy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesAddEnemy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
