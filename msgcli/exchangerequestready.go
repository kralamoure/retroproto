// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeRequestReady struct{}

func (m ExchangeRequestReady) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeRequestReady
}

func (m ExchangeRequestReady) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeRequestReady) Deserialize(extra string) error {
	return nil
}
