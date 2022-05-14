package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SpellsSpellForgetClose struct{}

func NewSpellsSpellForgetClose(extra string) (SpellsSpellForgetClose, error) {
	var m SpellsSpellForgetClose

	if err := m.Deserialize(extra); err != nil {
		return SpellsSpellForgetClose{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsSpellForgetClose) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellForgetClose
}

func (m SpellsSpellForgetClose) MessageName() string {
	return "SpellsSpellForgetClose"
}

func (m SpellsSpellForgetClose) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetClose) Deserialize(extra string) error {
	return nil
}
