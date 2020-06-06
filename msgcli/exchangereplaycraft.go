// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeReplayCraft struct{}

func (m ExchangeReplayCraft) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeReplayCraft
}

func (m ExchangeReplayCraft) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeReplayCraft) Deserialize(extra string) error {
	return nil
}
