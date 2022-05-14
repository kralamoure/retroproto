package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type GameCreate struct {
	Type int
}

func NewGameCreate(extra string) (GameCreate, error) {
	var m GameCreate

	if err := m.Deserialize(extra); err != nil {
		return GameCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameCreate) MessageId() retroproto.MsgCliId {
	return retroproto.GameCreate
}

func (m GameCreate) MessageName() string {
	return "GameCreate"
}

func (m GameCreate) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Type), nil
}

func (m *GameCreate) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	typeValue, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Type = int(typeValue)

	return nil
}
