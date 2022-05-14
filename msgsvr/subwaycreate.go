// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayCreate struct{}

func NewSubwayCreate(extra string) (SubwayCreate, error) {
	var m SubwayCreate

	if err := m.Deserialize(extra); err != nil {
		return SubwayCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.SubwayCreate
}

func (m SubwayCreate) MessageName() string {
	return "SubwayCreate"
}

func (m SubwayCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
