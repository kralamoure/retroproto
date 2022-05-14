// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildCreate struct{}

func NewGuildCreate(extra string) (GuildCreate, error) {
	var m GuildCreate

	if err := m.Deserialize(extra); err != nil {
		return GuildCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildCreate) MessageId() retroproto.MsgCliId {
	return retroproto.GuildCreate
}

func (m GuildCreate) MessageName() string {
	return "GuildCreate"
}

func (m GuildCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
