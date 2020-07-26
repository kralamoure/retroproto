// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLeaveError struct{}

func (m ExchangeLeaveError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLeaveError
}

func (m ExchangeLeaveError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeLeaveError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
