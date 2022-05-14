// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesRemoveEnemyError struct{}

func NewEnemiesRemoveEnemyError(extra string) (EnemiesRemoveEnemyError, error) {
	var m EnemiesRemoveEnemyError

	if err := m.Deserialize(extra); err != nil {
		return EnemiesRemoveEnemyError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesRemoveEnemyError) MessageId() retroproto.MsgSvrId {
	return retroproto.EnemiesRemoveEnemyError
}

func (m EnemiesRemoveEnemyError) MessageName() string {
	return "EnemiesRemoveEnemyError"
}

func (m EnemiesRemoveEnemyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesRemoveEnemyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
