// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildRequestLeave struct{}

func NewGuildRequestLeave(extra string) (GuildRequestLeave, error) {
	var m GuildRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return GuildRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.GuildRequestLeave
}

func (m GuildRequestLeave) MessageName() string {
	return "GuildRequestLeave"
}

func (m GuildRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
