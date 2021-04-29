package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type AccountNewLevel struct {
	Level int
}

func (m AccountNewLevel) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.AccountNewLevel
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
