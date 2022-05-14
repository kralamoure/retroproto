package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeLeaveError struct{}

func (m ExchangeLeaveError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLeaveError
}

func (m ExchangeLeaveError) MessageName() string {
	return "ExchangeLeaveError"
}

func (m ExchangeLeaveError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeLeaveError) Deserialize(extra string) error {
	return nil
}
