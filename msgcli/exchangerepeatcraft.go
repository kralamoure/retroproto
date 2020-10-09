// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeRepeatCraft struct{}

func (m ExchangeRepeatCraft) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeRepeatCraft
}

func (m ExchangeRepeatCraft) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeRepeatCraft) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
