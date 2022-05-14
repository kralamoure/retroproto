package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type GameActionAck struct {
	Id int
}

func NewGameActionAck(extra string) (GameActionAck, error) {
	var m GameActionAck

	if err := m.Deserialize(extra); err != nil {
		return GameActionAck{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameActionAck) MessageId() retroproto.MsgCliId {
	return retroproto.GameActionAck
}

func (m GameActionAck) MessageName() string {
	return "GameActionAck"
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
