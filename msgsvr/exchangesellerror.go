// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeSellError struct{}

func NewExchangeSellError(extra string) (ExchangeSellError, error) {
	var m ExchangeSellError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeSellError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeSellError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeSellError
}

func (m ExchangeSellError) MessageName() string {
	return "ExchangeSellError"
}

func (m ExchangeSellError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeSellError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
