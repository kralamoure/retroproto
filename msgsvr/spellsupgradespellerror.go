package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SpellsUpgradeSpellError struct{}

func NewSpellsUpgradeSpellError(extra string) (SpellsUpgradeSpellError, error) {
	var m SpellsUpgradeSpellError

	if err := m.Deserialize(extra); err != nil {
		return SpellsUpgradeSpellError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsUpgradeSpellError) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsUpgradeSpellError
}

func (m SpellsUpgradeSpellError) MessageName() string {
	return "SpellsUpgradeSpellError"
}

func (m SpellsUpgradeSpellError) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsUpgradeSpellError) Deserialize(extra string) error {
	return nil
}
