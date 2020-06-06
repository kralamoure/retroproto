// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeCreateSuccess struct{}

func (m ExchangeCreateSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeCreateSuccess
}

func (m ExchangeCreateSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeCreateSuccess) Deserialize(extra string) error {
	return nil
}
