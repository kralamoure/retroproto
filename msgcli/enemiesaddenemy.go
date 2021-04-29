// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesAddEnemy struct{}

func (m EnemiesAddEnemy) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.EnemiesAddEnemy
}

func (m EnemiesAddEnemy) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesAddEnemy) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
