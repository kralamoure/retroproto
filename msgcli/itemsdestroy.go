package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ItemsDestroy struct {
	Id       int
	Quantity int
}

func NewItemsDestroy(extra string) (ItemsDestroy, error) {
	var m ItemsDestroy

	if err := m.Deserialize(extra); err != nil {
		return ItemsDestroy{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsDestroy) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsDestroy
}

func (m ItemsDestroy) MessageName() string {
	return "ItemsDestroy"
}

func (m ItemsDestroy) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Quantity), nil
}

func (m *ItemsDestroy) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	quantity, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Quantity = int(quantity)

	return nil
}
