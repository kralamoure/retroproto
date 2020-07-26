// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesAddEnemy struct{}

func (m EnemiesAddEnemy) ProtocolId() d1proto.MsgCliId {
	return d1proto.EnemiesAddEnemy
}

func (m EnemiesAddEnemy) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *EnemiesAddEnemy) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
