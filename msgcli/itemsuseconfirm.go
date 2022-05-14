package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ItemsUseConfirm struct {
	Id       int
	SpriteId int
	Cell     int
}

func NewItemsUseConfirm(extra string) (ItemsUseConfirm, error) {
	var m ItemsUseConfirm

	if err := m.Deserialize(extra); err != nil {
		return ItemsUseConfirm{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsUseConfirm) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsUseConfirm
}

func (m ItemsUseConfirm) MessageName() string {
	return "ItemsUseConfirm"
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
