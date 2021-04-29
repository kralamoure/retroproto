package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
)

type ItemsUseConfirm struct {
	Id       int
	SpriteId int
	Cell     int
}

func (m ItemsUseConfirm) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ItemsUseConfirm
}

func (m ItemsUseConfirm) Serialized() (string, error) {
	spriteId := ""
	if m.SpriteId != 0 {
		spriteId = fmt.Sprintf("%d", m.SpriteId)
	}

	cell := ""
	if m.Cell != 0 {
		cell = fmt.Sprintf("|%d", m.Cell)
	}

	return fmt.Sprintf("%d|%s%s", m.Id, spriteId, cell), nil
}

func (m *ItemsUseConfirm) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	if len(sli) >= 2 && sli[1] != "" {
		spriteId, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.SpriteId = int(spriteId)
	}

	if len(sli) >= 3 && sli[2] != "" {
		cell, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Cell = int(cell)
	}

	return nil
}
