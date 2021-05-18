// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeRepeatCraft struct{}

func (m ExchangeRepeatCraft) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeRepeatCraft
}

func (m ExchangeRepeatCraft) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRepeatCraft) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
