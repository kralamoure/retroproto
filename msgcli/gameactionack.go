package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type GameActionAck struct {
	Id int
}

func (m GameActionAck) MessageId() retroproto.MsgCliId {
	return retroproto.GameActionAck
}

func (m GameActionAck) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *GameActionAck) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
