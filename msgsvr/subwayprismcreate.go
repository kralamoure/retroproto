// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayPrismCreate struct{}

func NewSubwayPrismCreate(extra string) (SubwayPrismCreate, error) {
	var m SubwayPrismCreate

	if err := m.Deserialize(extra); err != nil {
		return SubwayPrismCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayPrismCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.SubwayPrismCreate
}

func (m SubwayPrismCreate) MessageName() string {
	return "SubwayPrismCreate"
}

func (m SubwayPrismCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayPrismCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
