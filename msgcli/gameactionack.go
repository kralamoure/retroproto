package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type GameActionAck struct {
	Id int
}

func (m GameActionAck) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameActionAck
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
