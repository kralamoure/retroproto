// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesAddEnemyError struct{}

func NewEnemiesAddEnemyError(extra string) (EnemiesAddEnemyError, error) {
	var m EnemiesAddEnemyError

	if err := m.Deserialize(extra); err != nil {
		return EnemiesAddEnemyError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesAddEnemyError) MessageId() retroproto.MsgSvrId {
	return retroproto.EnemiesAddEnemyError
}

func (m EnemiesAddEnemyError) MessageName() string {
	return "EnemiesAddEnemyError"
}

func (m EnemiesAddEnemyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesAddEnemyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
