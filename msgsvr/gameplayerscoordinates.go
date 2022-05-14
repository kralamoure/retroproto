// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GamePlayersCoordinates struct{}

func NewGamePlayersCoordinates(extra string) (GamePlayersCoordinates, error) {
	var m GamePlayersCoordinates

	if err := m.Deserialize(extra); err != nil {
		return GamePlayersCoordinates{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GamePlayersCoordinates) MessageId() retroproto.MsgSvrId {
	return retroproto.GamePlayersCoordinates
}

func (m GamePlayersCoordinates) MessageName() string {
	return "GamePlayersCoordinates"
}

func (m GamePlayersCoordinates) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePlayersCoordinates) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
