package msgsvr

import (
	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/enum"
)

type ExchangeRequestError struct {
	Reason rune
}

func (m ExchangeRequestError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeRequestError
}

func (m ExchangeRequestError) Serialized() (string, error) {
	return string(m.Reason), nil
}

func (m *ExchangeRequestError) Deserialize(extra string) error {
	if extra == "" {
		m.Reason = enum.ExchangeRequestErrorReason.Default
	} else {
		m.Reason = rune(extra[0])
	}
	return nil
}
