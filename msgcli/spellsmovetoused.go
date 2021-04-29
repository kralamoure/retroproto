package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type SpellsMoveToUsed struct {
	Id       int
	Position int
}

func (m SpellsMoveToUsed) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.SpellsMoveToUsed
}

func (m SpellsMoveToUsed) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Position), nil
}

func (m *SpellsMoveToUsed) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return d1encoding.ErrInvalidMsg
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
