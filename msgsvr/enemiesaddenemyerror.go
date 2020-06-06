// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesAddEnemyError struct{}

func (m EnemiesAddEnemyError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EnemiesAddEnemyError
}

func (m EnemiesAddEnemyError) Serialized() (string, error) {
	return "", nil
}

func (m *EnemiesAddEnemyError) Deserialize(extra string) error {
	return nil
}
