// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLocalDistantError struct{}

func (m ExchangeLocalDistantError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLocalDistantError
}

func (m ExchangeLocalDistantError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeLocalDistantError) Deserialize(extra string) error {
	return nil
}
