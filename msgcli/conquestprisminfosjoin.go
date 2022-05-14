// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosJoin struct{}

func NewConquestPrismInfosJoin(extra string) (ConquestPrismInfosJoin, error) {
	var m ConquestPrismInfosJoin

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismInfosJoin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismInfosJoin) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestPrismInfosJoin
}

func (m ConquestPrismInfosJoin) MessageName() string {
	return "ConquestPrismInfosJoin"
}

func (m ConquestPrismInfosJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
