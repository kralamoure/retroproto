package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type GameCreateSuccess struct {
	Type int
}

func NewGameCreateSuccess(extra string) (GameCreateSuccess, error) {
	var m GameCreateSuccess

	if err := m.Deserialize(extra); err != nil {
		return GameCreateSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameCreateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GameCreateSuccess
}

func (m GameCreateSuccess) MessageName() string {
	return "GameCreateSuccess"
}

func (m GameCreateSuccess) Serialized() (string, error) {
	return fmt.Sprintf("|%d", m.Type), nil
}

func (m *GameCreateSuccess) Deserialize(extra string) error {
	extra = strings.TrimPrefix(extra, "|")
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	typeValue, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Type = int(typeValue)

	return nil
}
