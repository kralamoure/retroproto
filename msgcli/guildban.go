// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildBan struct{}

func NewGuildBan(extra string) (GuildBan, error) {
	var m GuildBan

	if err := m.Deserialize(extra); err != nil {
		return GuildBan{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildBan) MessageId() retroproto.MsgCliId {
	return retroproto.GuildBan
}

func (m GuildBan) MessageName() string {
	return "GuildBan"
}

func (m GuildBan) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildBan) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
