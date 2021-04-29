package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type GameMovementRemove struct {
	Id int
}

func (m GameMovementRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GameMovementRemove
}

func (m GameMovementRemove) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *GameMovementRemove) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1encoding.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
