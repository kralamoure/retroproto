// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildGetInfosMembers struct{}

func NewGuildGetInfosMembers(extra string) (GuildGetInfosMembers, error) {
	var m GuildGetInfosMembers

	if err := m.Deserialize(extra); err != nil {
		return GuildGetInfosMembers{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildGetInfosMembers) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosMembers
}

func (m GuildGetInfosMembers) MessageName() string {
	return "GuildGetInfosMembers"
}

func (m GuildGetInfosMembers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosMembers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
