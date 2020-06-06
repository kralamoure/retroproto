// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLeave struct{}

func (m ExchangeLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeLeave
}

func (m ExchangeLeave) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeLeave) Deserialize(extra string) error {
	return nil
}
