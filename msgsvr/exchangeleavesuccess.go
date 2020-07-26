// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeLeaveSuccess struct{}

func (m ExchangeLeaveSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeLeaveSuccess
}

func (m ExchangeLeaveSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangeLeaveSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
