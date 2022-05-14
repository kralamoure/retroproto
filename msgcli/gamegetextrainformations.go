package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameGetExtraInformations struct{}

func NewGameGetExtraInformations(extra string) (GameGetExtraInformations, error) {
	var m GameGetExtraInformations

	if err := m.Deserialize(extra); err != nil {
		return GameGetExtraInformations{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameGetExtraInformations) MessageId() retroproto.MsgCliId {
	return retroproto.GameGetExtraInformations
}

func (m GameGetExtraInformations) MessageName() string {
	return "GameGetExtraInformations"
}

func (m GameGetExtraInformations) Serialized() (string, error) {
	return "", nil
}

func (m *GameGetExtraInformations) Deserialize(extra string) error {
	return nil
}
