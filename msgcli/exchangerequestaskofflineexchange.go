// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeRequestAskOfflineExchange struct{}

func (m ExchangeRequestAskOfflineExchange) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeRequestAskOfflineExchange
}

func (m ExchangeRequestAskOfflineExchange) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeRequestAskOfflineExchange) Deserialize(extra string) error {
	return nil
}
