// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesSellError struct{}

func NewHousesSellError(extra string) (HousesSellError, error) {
	var m HousesSellError

	if err := m.Deserialize(extra); err != nil {
		return HousesSellError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesSellError) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesSellError
}

func (m HousesSellError) MessageName() string {
	return "HousesSellError"
}

func (m HousesSellError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesSellError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
