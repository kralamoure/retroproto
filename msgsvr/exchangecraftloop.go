// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeCraftLoop struct{}

func NewExchangeCraftLoop(extra string) (ExchangeCraftLoop, error) {
	var m ExchangeCraftLoop

	if err := m.Deserialize(extra); err != nil {
		return ExchangeCraftLoop{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeCraftLoop) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCraftLoop
}

func (m ExchangeCraftLoop) MessageName() string {
	return "ExchangeCraftLoop"
}

func (m ExchangeCraftLoop) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeCraftLoop) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
