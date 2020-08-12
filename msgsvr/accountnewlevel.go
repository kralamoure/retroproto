package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type AccountNewLevel struct {
	Level int
}

func (m AccountNewLevel) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountNewLevel
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
