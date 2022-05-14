// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildBoostSpell struct{}

func NewGuildBoostSpell(extra string) (GuildBoostSpell, error) {
	var m GuildBoostSpell

	if err := m.Deserialize(extra); err != nil {
		return GuildBoostSpell{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildBoostSpell) MessageId() retroproto.MsgCliId {
	return retroproto.GuildBoostSpell
}

func (m GuildBoostSpell) MessageName() string {
	return "GuildBoostSpell"
}

func (m GuildBoostSpell) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildBoostSpell) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
