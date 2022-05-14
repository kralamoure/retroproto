// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeSellSuccess struct{}

func NewExchangeSellSuccess(extra string) (ExchangeSellSuccess, error) {
	var m ExchangeSellSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeSellSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeSellSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeSellSuccess
}

func (m ExchangeSellSuccess) MessageName() string {
	return "ExchangeSellSuccess"
}

func (m ExchangeSellSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeSellSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
