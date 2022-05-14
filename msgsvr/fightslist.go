// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type FightsList struct{}

func NewFightsList(extra string) (FightsList, error) {
	var m FightsList

	if err := m.Deserialize(extra); err != nil {
		return FightsList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsList) MessageId() retroproto.MsgSvrId {
	return retroproto.FightsList
}

func (m FightsList) MessageName() string {
	return "FightsList"
}

func (m FightsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
