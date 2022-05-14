// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesBuySuccess struct{}

func NewHousesBuySuccess(extra string) (HousesBuySuccess, error) {
	var m HousesBuySuccess

	if err := m.Deserialize(extra); err != nil {
		return HousesBuySuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesBuySuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesBuySuccess
}

func (m HousesBuySuccess) MessageName() string {
	return "HousesBuySuccess"
}

func (m HousesBuySuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesBuySuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
