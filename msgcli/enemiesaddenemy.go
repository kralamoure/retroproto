// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesAddEnemy struct{}

func (m EnemiesAddEnemy) MessageId() retroproto.MsgCliId {
	return retroproto.EnemiesAddEnemy
}

func (m EnemiesAddEnemy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesAddEnemy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
