// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubwayUseError struct{}

func NewSubwayUseError(extra string) (SubwayUseError, error) {
	var m SubwayUseError

	if err := m.Deserialize(extra); err != nil {
		return SubwayUseError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubwayUseError) MessageId() retroproto.MsgSvrId {
	return retroproto.SubwayUseError
}

func (m SubwayUseError) MessageName() string {
	return "SubwayUseError"
}

func (m SubwayUseError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayUseError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
