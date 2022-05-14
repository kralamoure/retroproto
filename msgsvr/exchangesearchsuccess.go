package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeSearchSuccess struct{}

func NewExchangeSearchSuccess(extra string) (ExchangeSearchSuccess, error) {
	var m ExchangeSearchSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeSearchSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeSearchSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeSearchSuccess
}

func (m ExchangeSearchSuccess) MessageName() string {
	return "ExchangeSearchSuccess"
}

func (m ExchangeSearchSuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeSearchSuccess) Deserialize(extra string) error {
	return nil
}
