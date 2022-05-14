// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeLocalDistantSuccess struct{}

func NewExchangeLocalDistantSuccess(extra string) (ExchangeLocalDistantSuccess, error) {
	var m ExchangeLocalDistantSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeLocalDistantSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeLocalDistantSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeLocalDistantSuccess
}

func (m ExchangeLocalDistantSuccess) MessageName() string {
	return "ExchangeLocalDistantSuccess"
}

func (m ExchangeLocalDistantSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeLocalDistantSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
