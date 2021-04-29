package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1"

	"github.com/kralamoure/d1proto"
)

type SpellsList struct {
	Spells []d1.CharacterSpell
}

func (m SpellsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsList
}

func (m SpellsList) Serialized() (string, error) {
	spells := make([]string, len(m.Spells))
	for i, v := range m.Spells {
		position, err := d1proto.Encode64(v.Position)
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
	m.Spells = make([]d1.CharacterSpell, len(sli))
	for i, v := range sli {
		var spell d1.CharacterSpell

		sli2 := strings.Split(v, "~")
		if len(sli2) != 3 {
			return d1proto.ErrInvalidMsg
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
			position, err := d1proto.Decode64(v)
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
