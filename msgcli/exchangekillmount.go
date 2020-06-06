// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeKillMount struct{}

func (m ExchangeKillMount) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeKillMount
}

func (m ExchangeKillMount) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeKillMount) Deserialize(extra string) error {
	return nil
}
