// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeRequestSuccess struct{}

func (m ExchangeRequestSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeRequestSuccess
}

func (m ExchangeRequestSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeRequestSuccess) Deserialize(extra string) error {
	return nil
}
