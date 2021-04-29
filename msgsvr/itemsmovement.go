package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1encoding"
)

type ItemsMovement struct {
	Id       int
	Position d1typ.CharacterItemPosition
}

func (m ItemsMovement) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsMovement
}

func (m ItemsMovement) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Position), nil
}

func (m *ItemsMovement) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 1 {
		return d1encoding.ErrInvalidMsg
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
		m.Position = d1typ.CharacterItemPosition(position)
	}

	return nil
}
