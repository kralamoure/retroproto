// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildHireTaxCollectorSuccess struct{}

func NewGildHireTaxCollectorSuccess(extra string) (GildHireTaxCollectorSuccess, error) {
	var m GildHireTaxCollectorSuccess

	if err := m.Deserialize(extra); err != nil {
		return GildHireTaxCollectorSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildHireTaxCollectorSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GildHireTaxCollectorSuccess
}

func (m GildHireTaxCollectorSuccess) MessageName() string {
	return "GildHireTaxCollectorSuccess"
}

func (m GildHireTaxCollectorSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildHireTaxCollectorSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
