package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SpellsSpellForgetShow struct{}

func NewSpellsSpellForgetShow(extra string) (SpellsSpellForgetShow, error) {
	var m SpellsSpellForgetShow

	if err := m.Deserialize(extra); err != nil {
		return SpellsSpellForgetShow{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsSpellForgetShow) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsSpellForgetShow
}

func (m SpellsSpellForgetShow) MessageName() string {
	return "SpellsSpellForgetShow"
}

func (m SpellsSpellForgetShow) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsSpellForgetShow) Deserialize(extra string) error {
	return nil
}
