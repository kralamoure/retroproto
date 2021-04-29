// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesRemoveEnemy struct{}

func (m EnemiesRemoveEnemy) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.EnemiesRemoveEnemy
}

func (m EnemiesRemoveEnemy) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesRemoveEnemy) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
