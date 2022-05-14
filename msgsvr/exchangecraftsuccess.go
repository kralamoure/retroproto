// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCraftSuccess struct{}

func NewExchangeCraftSuccess(extra string) (ExchangeCraftSuccess, error) {
	var m ExchangeCraftSuccess

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCraftSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCraftSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftSuccess
}

func (m ExchangeCraftSuccess) MessageName() string {
	return "ExchangeCraftSuccess"
}

func (m ExchangeCraftSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
