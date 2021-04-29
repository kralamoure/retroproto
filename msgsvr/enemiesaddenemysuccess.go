// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesAddEnemySuccess struct{}

func (m EnemiesAddEnemySuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EnemiesAddEnemySuccess
}

func (m EnemiesAddEnemySuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesAddEnemySuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
