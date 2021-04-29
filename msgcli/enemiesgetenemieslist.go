// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type EnemiesGetEnemiesList struct{}

func (m EnemiesGetEnemiesList) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.EnemiesGetEnemiesList
}

func (m EnemiesGetEnemiesList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EnemiesGetEnemiesList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
