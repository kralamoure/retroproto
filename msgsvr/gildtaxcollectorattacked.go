// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildTaxCollectorAttacked struct{}

func NewGildTaxCollectorAttacked(extra string) (GildTaxCollectorAttacked, error) {
	var m GildTaxCollectorAttacked

	if err := m.Deserialize(extra); err != nil {
		return GildTaxCollectorAttacked{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildTaxCollectorAttacked) MessageId() retroproto.MsgSvrId {
	return retroproto.GildTaxCollectorAttacked
}

func (m GildTaxCollectorAttacked) MessageName() string {
	return "GildTaxCollectorAttacked"
}

func (m GildTaxCollectorAttacked) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildTaxCollectorAttacked) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
