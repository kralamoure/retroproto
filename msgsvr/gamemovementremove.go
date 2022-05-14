package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type GameMovementRemove struct {
	Id int
}

func (m GameMovementRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.GameMovementRemove
}

func (m GameMovementRemove) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *GameMovementRemove) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
