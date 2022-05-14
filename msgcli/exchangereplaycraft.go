// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeReplayCraft struct{}

func (m ExchangeReplayCraft) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeReplayCraft
}

func (m ExchangeReplayCraft) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeReplayCraft) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
