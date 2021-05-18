// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesRemoveEnemy struct{}

func (m EnemiesRemoveEnemy) ProtocolId() retroproto.MsgCliId {
	return retroproto.EnemiesRemoveEnemy
}

func (m EnemiesRemoveEnemy) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesRemoveEnemy) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
