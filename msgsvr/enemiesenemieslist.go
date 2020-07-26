// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesEnemiesList struct{}

func (m EnemiesEnemiesList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EnemiesEnemiesList
}

func (m EnemiesEnemiesList) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *EnemiesEnemiesList) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
