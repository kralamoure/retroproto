package msgsvr

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kralamoure/d1proto"
)

type BasicsTime struct {
	Value time.Time
}

func (m BasicsTime) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsTime
}

func (m BasicsTime) Serialized() (string, error) {
	valueN := m.Value.Unix()
	if valueN < 0 {
		return "", d1proto.ErrInvalidMsg
	}
	valueN *= 1000

	return fmt.Sprintf("%d", valueN), nil
}

func (m *BasicsTime) Deserialize(extra string) error {
	valueN, err := strconv.ParseInt(extra, 10, 64)
	if err != nil {
		return err
	}
	if valueN < 0 {
		return d1proto.ErrInvalidMsg
	}
	valueN /= 1000

	m.Value = time.Unix(valueN, 0)

	return nil
}
