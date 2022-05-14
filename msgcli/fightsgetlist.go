// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsGetList struct{}

func NewFightsGetList(extra string) (FightsGetList, error) {
	var m FightsGetList

	if err := m.Deserialize(extra); err != nil {
		return FightsGetList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsGetList) MessageId() retroproto.MsgCliId {
	return retroproto.FightsGetList
}

func (m FightsGetList) MessageName() string {
	return "FightsGetList"
}

func (m FightsGetList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsGetList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
