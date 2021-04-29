package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AccountQueuePosition struct{}

func (m AccountQueuePosition) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AccountQueuePosition
}

func (m AccountQueuePosition) Serialized() (string, error) {
	return "", nil
}

func (m *AccountQueuePosition) Deserialize(extra string) error {
	return nil
}
