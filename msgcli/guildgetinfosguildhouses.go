// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildGetInfosGuildHouses struct{}

func NewGuildGetInfosGuildHouses(extra string) (GuildGetInfosGuildHouses, error) {
	var m GuildGetInfosGuildHouses

	if err := m.Deserialize(extra); err != nil {
		return GuildGetInfosGuildHouses{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildGetInfosGuildHouses) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosGuildHouses
}

func (m GuildGetInfosGuildHouses) MessageName() string {
	return "GuildGetInfosGuildHouses"
}

func (m GuildGetInfosGuildHouses) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosGuildHouses) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
