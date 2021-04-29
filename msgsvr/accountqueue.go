package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type AccountQueue struct {
	Position int
}

func (m AccountQueue) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountQueue
}

func (m AccountQueue) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Position), nil
}

func (m *AccountQueue) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	position, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Position = int(position)

	return nil
}
