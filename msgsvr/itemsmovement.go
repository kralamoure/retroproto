package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type ItemsMovement struct {
	Id       int
	Position retrotyp.CharacterItemPosition
}

func (m ItemsMovement) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsMovement
}

func (m ItemsMovement) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Position), nil
}

func (m *ItemsMovement) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 1 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	if len(sli) >= 2 {
		position, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Position = retrotyp.CharacterItemPosition(position)
	}

	return nil
}
