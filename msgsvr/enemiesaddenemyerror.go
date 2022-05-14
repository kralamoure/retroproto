// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesAddEnemyError struct{}

func (m EnemiesAddEnemyError) MessageId() retroproto.MsgSvrId {
	return retroproto.EnemiesAddEnemyError
}

func (m EnemiesAddEnemyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesAddEnemyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
