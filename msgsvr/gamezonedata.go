// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameZoneData struct{}

func NewGameZoneData(extra string) (GameZoneData, error) {
	var m GameZoneData

	if err := m.Deserialize(extra); err != nil {
		return GameZoneData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameZoneData) MessageId() retroproto.MsgSvrId {
	return retroproto.GameZoneData
}

func (m GameZoneData) MessageName() string {
	return "GameZoneData"
}

func (m GameZoneData) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameZoneData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
