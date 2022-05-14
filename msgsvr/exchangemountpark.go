// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMountPark struct{}

func NewExchangeMountPark(extra string) (ExchangeMountPark, error) {
	var m ExchangeMountPark

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMountPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMountPark) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeMountPark
}

func (m ExchangeMountPark) MessageName() string {
	return "ExchangeMountPark"
}

func (m ExchangeMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
