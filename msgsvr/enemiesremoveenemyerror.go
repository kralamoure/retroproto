// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesRemoveEnemyError struct{}

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
