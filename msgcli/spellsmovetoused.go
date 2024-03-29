package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type SpellsMoveToUsed struct {
	Id       int
	Position int
}

func NewSpellsMoveToUsed(extra string) (SpellsMoveToUsed, error) {
	var m SpellsMoveToUsed

	if err := m.Deserialize(extra); err != nil {
		return SpellsMoveToUsed{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpellsMoveToUsed) MessageId() retroproto.MsgCliId {
	return retroproto.SpellsMoveToUsed
}

func (m SpellsMoveToUsed) MessageName() string {
	return "SpellsMoveToUsed"
}

func (m SpellsMoveToUsed) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Position), nil
}

func (m *SpellsMoveToUsed) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	position, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Position = int(position)

	return nil
}
