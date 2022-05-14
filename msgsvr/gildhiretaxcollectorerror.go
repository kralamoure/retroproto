// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildHireTaxCollectorError struct{}

func NewGildHireTaxCollectorError(extra string) (GildHireTaxCollectorError, error) {
	var m GildHireTaxCollectorError

	if err := m.Deserialize(extra); err != nil {
		return GildHireTaxCollectorError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildHireTaxCollectorError) MessageId() retroproto.MsgSvrId {
	return retroproto.GildHireTaxCollectorError
}

func (m GildHireTaxCollectorError) MessageName() string {
	return "GildHireTaxCollectorError"
}

func (m GildHireTaxCollectorError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildHireTaxCollectorError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
