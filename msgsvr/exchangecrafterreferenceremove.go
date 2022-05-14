// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCrafterReferenceRemove struct{}

func NewExchangeCrafterReferenceRemove(extra string) (ExchangeCrafterReferenceRemove, error) {
	var m ExchangeCrafterReferenceRemove

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCrafterReferenceRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCrafterReferenceRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCrafterReferenceRemove
}

func (m ExchangeCrafterReferenceRemove) MessageName() string {
	return "ExchangeCrafterReferenceRemove"
}

func (m ExchangeCrafterReferenceRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
