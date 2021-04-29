// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesAddEnemyError struct{}

func (m EnemiesAddEnemyError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EnemiesAddEnemyError
}

func (m EnemiesAddEnemyError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesAddEnemyError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
