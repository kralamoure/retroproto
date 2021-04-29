package msgsvr

import (
	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/enum"
)

type ExchangeRequestError struct {
	Reason rune
}

func (m ExchangeRequestError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeRequestError
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
