package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameCreateError struct{}

func NewGameCreateError(extra string) (GameCreateError, error) {
	var m GameCreateError

	if err := m.Deserialize(extra); err != nil {
		return GameCreateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCreateError
}

func (m GameCreateError) MessageName() string {
	return "GameCreateError"
}

func (m GameCreateError) Serialized() (string, error) {
	return "", nil
}

func (m *GameCreateError) Deserialize(extra string) error {
	return nil
}
