package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeBuySuccess struct{}

func NewExchangeBuySuccess(extra string) (ExchangeBuySuccess, error) {
	var m ExchangeBuySuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeBuySuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeBuySuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBuySuccess
}

func (m ExchangeBuySuccess) MessageName() string {
	return "ExchangeBuySuccess"
}

func (m ExchangeBuySuccess) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBuySuccess) Deserialize(extra string) error {
	return nil
}
