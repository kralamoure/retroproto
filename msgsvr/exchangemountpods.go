// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeMountPods struct{}

func NewExchangeMountPods(extra string) (ExchangeMountPods, error) {
	var m ExchangeMountPods

	if err := m.Deserialize(extra); err != nil {
		return ExchangeMountPods{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeMountPods) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeMountPods
}

func (m ExchangeMountPods) MessageName() string {
	return "ExchangeMountPods"
}

func (m ExchangeMountPods) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeMountPods) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
