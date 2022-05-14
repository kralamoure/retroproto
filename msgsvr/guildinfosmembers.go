// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosMembers struct{}

func NewGuildInfosMembers(extra string) (GuildInfosMembers, error) {
	var m GuildInfosMembers

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosMembers{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosMembers) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosMembers
}

func (m GuildInfosMembers) MessageName() string {
	return "GuildInfosMembers"
}

func (m GuildInfosMembers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosMembers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
