package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type InfosLifeRestoreTimerFinish struct {
	Restored int
}

func NewInfosLifeRestoreTimerFinish(extra string) (InfosLifeRestoreTimerFinish, error) {
	var m InfosLifeRestoreTimerFinish

	if err := m.Deserialize(extra); err != nil {
		return InfosLifeRestoreTimerFinish{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosLifeRestoreTimerFinish) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosLifeRestoreTimerFinish
}

func (m InfosLifeRestoreTimerFinish) MessageName() string {
	return "InfosLifeRestoreTimerFinish"
}

func (m InfosLifeRestoreTimerFinish) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Restored), nil
}

func (m *InfosLifeRestoreTimerFinish) Deserialize(extra string) error {
	if len(extra) < 1 {
		return nil
	}

	restored, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Restored = int(restored)

	return nil
}
