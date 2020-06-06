// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeReady struct{}

func (m ExchangeReady) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeReady
}

func (m ExchangeReady) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeReady) Deserialize(extra string) error {
	return nil
}
