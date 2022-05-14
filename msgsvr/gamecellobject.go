// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameCellObject struct{}

func NewGameCellObject(extra string) (GameCellObject, error) {
	var m GameCellObject

	if err := m.Deserialize(extra); err != nil {
		return GameCellObject{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameCellObject) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCellObject
}

func (m GameCellObject) MessageName() string {
	return "GameCellObject"
}

func (m GameCellObject) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameCellObject) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
