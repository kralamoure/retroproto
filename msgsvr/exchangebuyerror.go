package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeBuyError struct{}

func NewExchangeBuyError(extra string) (ExchangeBuyError, error) {
	var m ExchangeBuyError

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBuyError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBuyError) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBuyError
}

func (m ExchangeBuyError) MessageName() string {
	return "ExchangeBuyError"
}

func (m ExchangeBuyError) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuyError) Deserialize(extra string) error {
	return nil
}
