package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameMapLoaded struct{}

func NewGameMapLoaded(extra string) (GameMapLoaded, error) {
	var m GameMapLoaded

	if err := m.Deserialize(extra); err != nil {
		return GameMapLoaded{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameMapLoaded) MessageId() retroproto.MsgSvrId {
	return retroproto.GameMapLoaded
}

func (m GameMapLoaded) MessageName() string {
	return "GameMapLoaded"
}

func (m GameMapLoaded) Serialized() (string, error) {
	return "", nil
}

func (m *GameMapLoaded) Deserialize(extra string) error {
	return nil
}
