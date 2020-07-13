package msgsvr

import (
	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/enum"
)

type ExchangeRequestError struct {
	Reason rune
}

func (m ExchangeRequestError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeRequestError
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
