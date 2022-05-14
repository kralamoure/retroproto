// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameRequestReady struct{}

func NewGameRequestReady(extra string) (GameRequestReady, error) {
	var m GameRequestReady

	if err := m.Deserialize(extra); err != nil {
		return GameRequestReady{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameRequestReady) MessageId() retroproto.MsgCliId {
	return retroproto.GameRequestReady
}

func (m GameRequestReady) MessageName() string {
	return "GameRequestReady"
}

func (m GameRequestReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameRequestReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
