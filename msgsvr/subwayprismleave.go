// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayPrismLeave struct{}

func NewSubwayPrismLeave(extra string) (SubwayPrismLeave, error) {
	var m SubwayPrismLeave

	if err := m.Deserialize(extra); err != nil {
		return SubwayPrismLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayPrismLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.SubwayPrismLeave
}

func (m SubwayPrismLeave) MessageName() string {
	return "SubwayPrismLeave"
}

func (m SubwayPrismLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
