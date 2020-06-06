// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesRemoveEnemyError struct{}

func (m EnemiesRemoveEnemyError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EnemiesRemoveEnemyError
}

func (m EnemiesRemoveEnemyError) Serialized() (string, error) {
	return "", nil
}

func (m *EnemiesRemoveEnemyError) Deserialize(extra string) error {
	return nil
}
