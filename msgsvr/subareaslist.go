// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type SubareasList struct{}

func NewSubareasList(extra string) (SubareasList, error) {
	var m SubareasList

	if err := m.Deserialize(extra); err != nil {
		return SubareasList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SubareasList) MessageId() retroproto.MsgSvrId {
	return retroproto.SubareasList
}

func (m SubareasList) MessageName() string {
	return "SubareasList"
}

func (m SubareasList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubareasList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
