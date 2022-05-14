// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameGetMapData struct{}

func NewGameGetMapData(extra string) (GameGetMapData, error) {
	var m GameGetMapData

	if err := m.Deserialize(extra); err != nil {
		return GameGetMapData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameGetMapData) MessageId() retroproto.MsgCliId {
	return retroproto.GameGetMapData
}

func (m GameGetMapData) MessageName() string {
	return "GameGetMapData"
}

func (m GameGetMapData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameGetMapData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
