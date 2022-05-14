// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosBoosts struct{}

func NewGuildInfosBoosts(extra string) (GuildInfosBoosts, error) {
	var m GuildInfosBoosts

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosBoosts{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosBoosts) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosBoosts
}

func (m GuildInfosBoosts) MessageName() string {
	return "GuildInfosBoosts"
}

func (m GuildInfosBoosts) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosBoosts) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
