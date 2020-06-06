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
	return "", nil
}

func (m *ExchangeLeaveSuccess) Deserialize(extra string) error {
	return nil
}
