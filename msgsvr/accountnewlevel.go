package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type AccountNewLevel struct {
	Level int
}

func NewAccountNewLevel(extra string) (AccountNewLevel, error) {
	var m AccountNewLevel

	if err := m.Deserialize(extra); err != nil {
		return AccountNewLevel{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountNewLevel) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountNewLevel
}

func (m AccountNewLevel) MessageName() string {
	return "AccountNewLevel"
}

func (m AccountNewLevel) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Level), nil
}

func (m *AccountNewLevel) Deserialize(extra string) error {
	level, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Level = int(level)

	return nil
}
