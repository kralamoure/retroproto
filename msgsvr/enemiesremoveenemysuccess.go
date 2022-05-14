// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesRemoveEnemySuccess struct{}

func NewEnemiesRemoveEnemySuccess(extra string) (EnemiesRemoveEnemySuccess, error) {
	var m EnemiesRemoveEnemySuccess

	if err := m.Deserialize(extra); err != nil {
		return EnemiesRemoveEnemySuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesRemoveEnemySuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.EnemiesRemoveEnemySuccess
}

func (m EnemiesRemoveEnemySuccess) MessageName() string {
	return "EnemiesRemoveEnemySuccess"
}

func (m EnemiesRemoveEnemySuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesRemoveEnemySuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
