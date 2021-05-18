package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLeaveError struct{}

func (m ExchangeLeaveError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ExchangeLeaveError
}

func (m ExchangeLeaveError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeLeaveError) Deserialize(extra string) error {
	return nil
}
