// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayLeave struct{}

func NewSubwayLeave(extra string) (SubwayLeave, error) {
	var m SubwayLeave

	if err := m.Deserialize(extra); err != nil {
		return SubwayLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.SubwayLeave
}

func (m SubwayLeave) MessageName() string {
	return "SubwayLeave"
}

func (m SubwayLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
