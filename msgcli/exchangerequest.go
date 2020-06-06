// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeRequest struct{}

func (m ExchangeRequest) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeRequest
}

func (m ExchangeRequest) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeRequest) Deserialize(extra string) error {
	return nil
}
