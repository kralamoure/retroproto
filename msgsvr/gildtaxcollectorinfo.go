// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildTaxCollectorInfo struct{}

func NewGildTaxCollectorInfo(extra string) (GildTaxCollectorInfo, error) {
	var m GildTaxCollectorInfo

	if err := m.Deserialize(extra); err != nil {
		return GildTaxCollectorInfo{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildTaxCollectorInfo) MessageId() retroproto.MsgSvrId {
	return retroproto.GildTaxCollectorInfo
}

func (m GildTaxCollectorInfo) MessageName() string {
	return "GildTaxCollectorInfo"
}

func (m GildTaxCollectorInfo) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildTaxCollectorInfo) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
