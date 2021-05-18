package msgsvr

import (
	"fmt"
	"time"

	"github.com/kralamoure/retroproto"
)

type InfosLifeRestoreTimerStart struct {
	Interval time.Duration
}

func (m InfosLifeRestoreTimerStart) ProtocolId() retroproto.MsgSvrId {
	return retroproto.InfosLifeRestoreTimerStart
}

func (m InfosLifeRestoreTimerStart) Serialized() (string, error) {
	interval := m.Interval.Milliseconds()
	if interval < 0 {
		interval = 0
	}

	return fmt.Sprintf("%d", interval), nil
}

func (m *InfosLifeRestoreTimerStart) Deserialize(extra string) error {
	if len(extra) < 1 {
		return nil
	}

	interval, err := time.ParseDuration(fmt.Sprintf("%sms", extra))
	if err != nil {
		return err
	}
	m.Interval = interval

	return nil
}
