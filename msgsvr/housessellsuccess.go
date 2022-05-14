// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesSellSuccess struct{}

func NewHousesSellSuccess(extra string) (HousesSellSuccess, error) {
	var m HousesSellSuccess

	if err := m.Deserialize(extra); err != nil {
		return HousesSellSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesSellSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesSellSuccess
}

func (m HousesSellSuccess) MessageName() string {
	return "HousesSellSuccess"
}

func (m HousesSellSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesSellSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
