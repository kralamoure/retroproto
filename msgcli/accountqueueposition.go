package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AccountQueuePosition struct{}

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
