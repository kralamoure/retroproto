package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AccountQueuePosition struct{}

func (m AccountQueuePosition) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountQueuePosition
}

func (m AccountQueuePosition) Serialized() (string, error) {
	return "", nil
}

func (m *AccountQueuePosition) Deserialize(extra string) error {
	return nil
}
