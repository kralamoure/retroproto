// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesEnemiesList struct{}

func (m EnemiesEnemiesList) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EnemiesEnemiesList
}

func (m EnemiesEnemiesList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesEnemiesList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
