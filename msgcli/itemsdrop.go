package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ItemsDrop struct {
	Id       int
	Quantity int
}

func (m ItemsDrop) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsDrop
}

func (m ItemsDrop) MessageName() string {
	return "ItemsDrop"
}

func (m ItemsDrop) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Quantity), nil
}

func (m *ItemsDrop) Deserialize(extra string) error {
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
