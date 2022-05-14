// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type EnemiesGetEnemiesList struct{}

func (m EnemiesGetEnemiesList) MessageId() retroproto.MsgCliId {
	return retroproto.EnemiesGetEnemiesList
}

func (m EnemiesGetEnemiesList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesGetEnemiesList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
