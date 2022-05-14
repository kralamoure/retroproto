// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCraftLoopEnd struct{}

func NewExchangeCraftLoopEnd(extra string) (ExchangeCraftLoopEnd, error) {
	var m ExchangeCraftLoopEnd

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCraftLoopEnd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCraftLoopEnd) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftLoopEnd
}

func (m ExchangeCraftLoopEnd) MessageName() string {
	return "ExchangeCraftLoopEnd"
}

func (m ExchangeCraftLoopEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftLoopEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
