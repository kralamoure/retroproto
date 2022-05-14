package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro"

	"github.com/kralamoure/retroproto"
)

type SpellsList struct {
	Spells []retro.CharacterSpell
}

func NewSpellsList(extra string) (SpellsList, error) {
	var m SpellsList

	if err := m.Deserialize(extra); err != nil {
		return SpellsList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsList) MessageId() retroproto.MsgSvrId {
	return retroproto.SpellsList
}

func (m SpellsList) MessageName() string {
	return "SpellsList"
}

func (m SpellsList) Serialized() (string, error) {
	spells := make([]string, len(m.Spells))
	for i, v := range m.Spells {
		position, err := retroproto.Encode64(v.Position)
		if err != nil {
			return "", err
		}
		spells[i] = fmt.Sprintf("%d~%d~%s", v.Id, v.Level, string(position))
	}

	return strings.Join(spells, ";"), nil
}

func (m *SpellsList) Deserialize(extra string) error {
	extra = strings.TrimSuffix(extra, ";")

	if extra == "" {
		return nil
	}

	sli := strings.Split(extra, ";")
	m.Spells = make([]retro.CharacterSpell, len(sli))
	for i, v := range sli {
		var spell retro.CharacterSpell

		sli2 := strings.Split(v, "~")
		if len(sli2) != 3 {
			return retroproto.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(sli2[0], 10, 32)
		if err != nil {
			return err
		}
		spell.Id = int(id)

		level, err := strconv.ParseInt(sli2[1], 10, 32)
		if err != nil {
			return err
		}
		spell.Level = int(level)

		for _, v := range sli2[2] {
			position, err := retroproto.Decode64(v)
			if err != nil {
				return err
			}
			spell.Position = position
			break
		}

		m.Spells[i] = spell
	}

	return nil
}
