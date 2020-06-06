// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeOfflineExchange struct{}

func (m ExchangeOfflineExchange) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeOfflineExchange
}

func (m ExchangeOfflineExchange) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeOfflineExchange) Deserialize(extra string) error {
	return nil
}
