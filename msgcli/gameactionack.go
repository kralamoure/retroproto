package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type GameActionAck struct {
	Id int
}

func (m GameActionAck) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameActionAck
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
