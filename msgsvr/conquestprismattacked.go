// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismAttacked struct{}

func NewConquestPrismAttacked(extra string) (ConquestPrismAttacked, error) {
	var m ConquestPrismAttacked

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismAttacked{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismAttacked) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismAttacked
}

func (m ConquestPrismAttacked) MessageName() string {
	return "ConquestPrismAttacked"
}

func (m ConquestPrismAttacked) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismAttacked) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
