package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type ItemsRequestMovement struct {
	Id       int
	Position retrotyp.CharacterItemPosition
	Quantity int
}

func (m ItemsRequestMovement) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsRequestMovement
}

func (m ItemsRequestMovement) MessageName() string {
	return "ItemsRequestMovement"
}

func (m ItemsRequestMovement) Serialized() (string, error) {
	quantity := ""
	if m.Quantity != 1 {
		quantity = fmt.Sprintf("|%d", m.Quantity)
	}
	return fmt.Sprintf("%d|%d%s", m.Id, m.Position, quantity), nil
}

func (m *ItemsRequestMovement) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
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
	m.Position = retrotyp.CharacterItemPosition(position)

	if len(sli) >= 3 {
		quantity, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		if quantity <= 0 {
			return retroproto.ErrInvalidMsg
		}
		m.Quantity = int(quantity)
	} else {
		m.Quantity = 1
	}

	return nil
}
