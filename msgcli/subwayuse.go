// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayUse struct{}

func NewSubwayUse(extra string) (SubwayUse, error) {
	var m SubwayUse

	if err := m.Deserialize(extra); err != nil {
		return SubwayUse{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayUse) MessageId() retroproto.MsgCliId {
	return retroproto.SubwayUse
}

func (m SubwayUse) MessageName() string {
	return "SubwayUse"
}

func (m SubwayUse) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayUse) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
