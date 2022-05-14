// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type InfosQuantity struct{}

func NewInfosQuantity(extra string) (InfosQuantity, error) {
	var m InfosQuantity

	if err := m.Deserialize(extra); err != nil {
		return InfosQuantity{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosQuantity) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosQuantity
}

func (m InfosQuantity) MessageName() string {
	return "InfosQuantity"
}

func (m InfosQuantity) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosQuantity) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
