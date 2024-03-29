package msgsvr

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kralamoure/retroproto"
)

type BasicsTime struct {
	Value time.Time
}

func NewBasicsTime(extra string) (BasicsTime, error) {
	var m BasicsTime

	if err := m.Deserialize(extra); err != nil {
		return BasicsTime{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsTime) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsTime
}

func (m BasicsTime) MessageName() string {
	return "BasicsTime"
}

func (m BasicsTime) Serialized() (string, error) {
	valueN := m.Value.Unix()
	if valueN < 0 {
		return "", retroproto.ErrInvalidMsg
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
		return retroproto.ErrInvalidMsg
	}
	valueN /= 1000

	m.Value = time.Unix(valueN, 0)

	return nil
}
