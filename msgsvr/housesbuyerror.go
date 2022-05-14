// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesBuyError struct{}

func NewHousesBuyError(extra string) (HousesBuyError, error) {
	var m HousesBuyError

	if err := m.Deserialize(extra); err != nil {
		return HousesBuyError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesBuyError) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesBuyError
}

func (m HousesBuyError) MessageName() string {
	return "HousesBuyError"
}

func (m HousesBuyError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesBuyError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
