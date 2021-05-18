package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountQueue struct {
	Position int
}

func (m AccountQueue) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountQueue
}

func (m AccountQueue) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Position), nil
}

func (m *AccountQueue) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	position, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Position = int(position)

	return nil
}
