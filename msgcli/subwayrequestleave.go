// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayRequestLeave struct{}

func NewSubwayRequestLeave(extra string) (SubwayRequestLeave, error) {
	var m SubwayRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return SubwayRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.SubwayRequestLeave
}

func (m SubwayRequestLeave) MessageName() string {
	return "SubwayRequestLeave"
}

func (m SubwayRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
