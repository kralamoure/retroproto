// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildLeave struct{}

func NewGuildLeave(extra string) (GuildLeave, error) {
	var m GuildLeave

	if err := m.Deserialize(extra); err != nil {
		return GuildLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildLeave
}

func (m GuildLeave) MessageName() string {
	return "GuildLeave"
}

func (m GuildLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
