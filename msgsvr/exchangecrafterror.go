// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCraftError struct{}

func NewExchangeCraftError(extra string) (ExchangeCraftError, error) {
	var m ExchangeCraftError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCraftError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCraftError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftError
}

func (m ExchangeCraftError) MessageName() string {
	return "ExchangeCraftError"
}

func (m ExchangeCraftError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
