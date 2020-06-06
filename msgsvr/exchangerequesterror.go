// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeRequestError struct{}

func (m ExchangeRequestError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeRequestError
}

func (m ExchangeRequestError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeRequestError) Deserialize(extra string) error {
	return nil
}
