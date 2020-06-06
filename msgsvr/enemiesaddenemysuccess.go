// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesAddEnemySuccess struct{}

func (m EnemiesAddEnemySuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EnemiesAddEnemySuccess
}

func (m EnemiesAddEnemySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *EnemiesAddEnemySuccess) Deserialize(extra string) error {
	return nil
}
