// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeSearchSuccess struct{}

func (m ExchangeSearchSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeSearchSuccess
}

func (m ExchangeSearchSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchSuccess) Deserialize(extra string) error {
	return nil
}
