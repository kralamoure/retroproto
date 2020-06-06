// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesRemoveEnemySuccess struct{}

func (m EnemiesRemoveEnemySuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EnemiesRemoveEnemySuccess
}

func (m EnemiesRemoveEnemySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *EnemiesRemoveEnemySuccess) Deserialize(extra string) error {
	return nil
}
