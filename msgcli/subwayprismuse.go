// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayPrismUse struct{}

func NewSubwayPrismUse(extra string) (SubwayPrismUse, error) {
	var m SubwayPrismUse

	if err := m.Deserialize(extra); err != nil {
		return SubwayPrismUse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayPrismUse) MessageId() retroproto.MsgCliId {
	return retroproto.SubwayPrismUse
}

func (m SubwayPrismUse) MessageName() string {
	return "SubwayPrismUse"
}

func (m SubwayPrismUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
