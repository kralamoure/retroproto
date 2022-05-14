// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismDead struct{}

func NewConquestPrismDead(extra string) (ConquestPrismDead, error) {
	var m ConquestPrismDead

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismDead{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismDead) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismDead
}

func (m ConquestPrismDead) MessageName() string {
	return "ConquestPrismDead"
}

func (m ConquestPrismDead) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismDead) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
