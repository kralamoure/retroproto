// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsNeedHelp struct{}

func NewFightsNeedHelp(extra string) (FightsNeedHelp, error) {
	var m FightsNeedHelp

	if err := m.Deserialize(extra); err != nil {
		return FightsNeedHelp{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsNeedHelp) MessageId() retroproto.MsgCliId {
	return retroproto.FightsNeedHelp
}

func (m FightsNeedHelp) MessageName() string {
	return "FightsNeedHelp"
}

func (m FightsNeedHelp) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsNeedHelp) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
