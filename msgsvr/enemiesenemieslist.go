// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesEnemiesList struct{}

func NewEnemiesEnemiesList(extra string) (EnemiesEnemiesList, error) {
	var m EnemiesEnemiesList

	if err := m.Deserialize(extra); err != nil {
		return EnemiesEnemiesList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesEnemiesList) MessageId() retroproto.MsgSvrId {
	return retroproto.EnemiesEnemiesList
}

func (m EnemiesEnemiesList) MessageName() string {
	return "EnemiesEnemiesList"
}

func (m EnemiesEnemiesList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesEnemiesList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
