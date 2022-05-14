// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosLeave struct{}

func NewConquestPrismInfosLeave(extra string) (ConquestPrismInfosLeave, error) {
	var m ConquestPrismInfosLeave

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismInfosLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismInfosLeave) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestPrismInfosLeave
}

func (m ConquestPrismInfosLeave) MessageName() string {
	return "ConquestPrismInfosLeave"
}

func (m ConquestPrismInfosLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
