// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesAddEnemySuccess struct{}

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
