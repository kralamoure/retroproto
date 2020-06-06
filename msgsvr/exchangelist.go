// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeList struct{}

func (m ExchangeList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeList
}

func (m ExchangeList) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeList) Deserialize(extra string) error {
	return nil
}
