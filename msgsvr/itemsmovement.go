package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ItemsMovement struct {
	Id       int
	Position int
}

func (m ItemsMovement) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsMovement
}

func (m ItemsMovement) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Position), nil
}

func (m *ItemsMovement) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 1 {
		return d1proto.ErrInvalidMsg
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
		m.Position = int(position)
	}

	return nil
}
