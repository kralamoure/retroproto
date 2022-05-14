// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesAddEnemySuccess struct{}

func NewEnemiesAddEnemySuccess(extra string) (EnemiesAddEnemySuccess, error) {
	var m EnemiesAddEnemySuccess

	if err := m.Deserialize(extra); err != nil {
		return EnemiesAddEnemySuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesAddEnemySuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.EnemiesAddEnemySuccess
}

func (m EnemiesAddEnemySuccess) MessageName() string {
	return "EnemiesAddEnemySuccess"
}

func (m EnemiesAddEnemySuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesAddEnemySuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
