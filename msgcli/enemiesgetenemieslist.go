// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type EnemiesGetEnemiesList struct{}

func NewEnemiesGetEnemiesList(extra string) (EnemiesGetEnemiesList, error) {
	var m EnemiesGetEnemiesList

	if err := m.Deserialize(extra); err != nil {
		return EnemiesGetEnemiesList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m EnemiesGetEnemiesList) MessageId() retroproto.MsgCliId {
	return retroproto.EnemiesGetEnemiesList
}

func (m EnemiesGetEnemiesList) MessageName() string {
	return "EnemiesGetEnemiesList"
}

func (m EnemiesGetEnemiesList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EnemiesGetEnemiesList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
