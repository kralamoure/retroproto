package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ItemsRequestMovement struct {
	Id       int
	Position int
	Quantity int
}

func (m ItemsRequestMovement) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsRequestMovement
}

func (m ItemsRequestMovement) Serialized() (string, error) {
	quantity := ""
	if m.Quantity != 0 {
		quantity = fmt.Sprintf("|%d", m.Quantity)
	}
	return fmt.Sprintf("%d|%d%s", m.Id, m.Position, quantity), nil
}

func (m *ItemsRequestMovement) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return d1proto.ErrInvalidMsg
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

	if len(sli) >= 3 {
		quantity, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Quantity = int(quantity)
	}

	return nil
}
