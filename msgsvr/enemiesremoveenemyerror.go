// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesRemoveEnemyError struct{}

func (m EnemiesRemoveEnemyError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EnemiesRemoveEnemyError
}

func (m EnemiesRemoveEnemyError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesRemoveEnemyError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
