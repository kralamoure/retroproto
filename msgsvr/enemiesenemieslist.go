// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesEnemiesList struct{}

func (m EnemiesEnemiesList) ProtocolId() retroproto.MsgSvrId {
	return retroproto.EnemiesEnemiesList
}

func (m EnemiesEnemiesList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesEnemiesList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
