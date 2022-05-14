package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AccountQueuePosition struct{}

func NewAccountQueuePosition(extra string) (AccountQueuePosition, error) {
	var m AccountQueuePosition

	if err := m.Deserialize(extra); err != nil {
		return AccountQueuePosition{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountQueuePosition) MessageId() retroproto.MsgCliId {
	return retroproto.AccountQueuePosition
}

func (m AccountQueuePosition) MessageName() string {
	return "AccountQueuePosition"
}

func (m AccountQueuePosition) Serialized() (string, error) {
	return "", nil
}

func (m *AccountQueuePosition) Deserialize(extra string) error {
	return nil
}
