// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesRemoveEnemySuccess struct{}

func (m EnemiesRemoveEnemySuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EnemiesRemoveEnemySuccess
}

func (m EnemiesRemoveEnemySuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesRemoveEnemySuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
