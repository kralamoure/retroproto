// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameCellData struct{}

func NewGameCellData(extra string) (GameCellData, error) {
	var m GameCellData

	if err := m.Deserialize(extra); err != nil {
		return GameCellData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameCellData) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCellData
}

func (m GameCellData) MessageName() string {
	return "GameCellData"
}

func (m GameCellData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameCellData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
