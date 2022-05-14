// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildStats struct{}

func NewGuildStats(extra string) (GuildStats, error) {
	var m GuildStats

	if err := m.Deserialize(extra); err != nil {
		return GuildStats{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildStats) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildStats
}

func (m GuildStats) MessageName() string {
	return "GuildStats"
}

func (m GuildStats) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildStats) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
