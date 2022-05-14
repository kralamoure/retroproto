// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayRequestPrismLeave struct{}

func NewSubwayRequestPrismLeave(extra string) (SubwayRequestPrismLeave, error) {
	var m SubwayRequestPrismLeave

	if err := m.Deserialize(extra); err != nil {
		return SubwayRequestPrismLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayRequestPrismLeave) MessageId() retroproto.MsgCliId {
	return retroproto.SubwayRequestPrismLeave
}

func (m SubwayRequestPrismLeave) MessageName() string {
	return "SubwayRequestPrismLeave"
}

func (m SubwayRequestPrismLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayRequestPrismLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
