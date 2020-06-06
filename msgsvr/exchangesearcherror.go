// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeSearchError struct{}

func (m ExchangeSearchError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeSearchError
}

func (m ExchangeSearchError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchError) Deserialize(extra string) error {
	return nil
}
