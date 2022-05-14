// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildGetInfosBoosts struct{}

func NewGuildGetInfosBoosts(extra string) (GuildGetInfosBoosts, error) {
	var m GuildGetInfosBoosts

	if err := m.Deserialize(extra); err != nil {
		return GuildGetInfosBoosts{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildGetInfosBoosts) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosBoosts
}

func (m GuildGetInfosBoosts) MessageName() string {
	return "GuildGetInfosBoosts"
}

func (m GuildGetInfosBoosts) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosBoosts) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
