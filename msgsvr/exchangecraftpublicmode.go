// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCraftPublicMode struct{}

func NewExchangeCraftPublicMode(extra string) (ExchangeCraftPublicMode, error) {
	var m ExchangeCraftPublicMode

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCraftPublicMode{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCraftPublicMode) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftPublicMode
}

func (m ExchangeCraftPublicMode) MessageName() string {
	return "ExchangeCraftPublicMode"
}

func (m ExchangeCraftPublicMode) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftPublicMode) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
