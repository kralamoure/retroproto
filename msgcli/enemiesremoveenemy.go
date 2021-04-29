// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesRemoveEnemy struct{}

func (m EnemiesRemoveEnemy) ProtocolId() d1proto.MsgCliId {
	return d1proto.EnemiesRemoveEnemy
}

func (m EnemiesRemoveEnemy) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *EnemiesRemoveEnemy) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
