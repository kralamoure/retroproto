package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type SpellsBoost struct {
	Id int
}

func NewSpellsBoost(extra string) (SpellsBoost, error) {
	var m SpellsBoost

	if err := m.Deserialize(extra); err != nil {
		return SpellsBoost{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsBoost) MessageId() retroproto.MsgCliId {
	return retroproto.SpellsBoost
}

func (m SpellsBoost) MessageName() string {
	return "SpellsBoost"
}

func (m SpellsBoost) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *SpellsBoost) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
