// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCrafterReferenceAdd struct{}

func NewExchangeCrafterReferenceAdd(extra string) (ExchangeCrafterReferenceAdd, error) {
	var m ExchangeCrafterReferenceAdd

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCrafterReferenceAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCrafterReferenceAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCrafterReferenceAdd
}

func (m ExchangeCrafterReferenceAdd) MessageName() string {
	return "ExchangeCrafterReferenceAdd"
}

func (m ExchangeCrafterReferenceAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCrafterReferenceAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
