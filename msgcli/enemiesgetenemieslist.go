// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type EnemiesGetEnemiesList struct{}

func (m EnemiesGetEnemiesList) ProtocolId() d1proto.MsgCliId {
	return d1proto.EnemiesGetEnemiesList
}

func (m EnemiesGetEnemiesList) Serialized() (string, error) {
	return "", nil
}

func (m *EnemiesGetEnemiesList) Deserialize(extra string) error {
	return nil
}
