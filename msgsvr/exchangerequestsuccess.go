// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeRequestSuccess struct{}

func NewExchangeRequestSuccess(extra string) (ExchangeRequestSuccess, error) {
	var m ExchangeRequestSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeRequestSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeRequestSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeRequestSuccess
}

func (m ExchangeRequestSuccess) MessageName() string {
	return "ExchangeRequestSuccess"
}

func (m ExchangeRequestSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRequestSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
