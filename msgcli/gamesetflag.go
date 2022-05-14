// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameSetFlag struct{}

func NewGameSetFlag(extra string) (GameSetFlag, error) {
	var m GameSetFlag

	if err := m.Deserialize(extra); err != nil {
		return GameSetFlag{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameSetFlag) MessageId() retroproto.MsgCliId {
	return retroproto.GameSetFlag
}

func (m GameSetFlag) MessageName() string {
	return "GameSetFlag"
}

func (m GameSetFlag) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameSetFlag) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
