// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesRemoveEnemySuccess struct{}

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
