// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosHouses struct{}

func NewGuildInfosHouses(extra string) (GuildInfosHouses, error) {
	var m GuildInfosHouses

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosHouses{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosHouses) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosHouses
}

func (m GuildInfosHouses) MessageName() string {
	return "GuildInfosHouses"
}

func (m GuildInfosHouses) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosHouses) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
