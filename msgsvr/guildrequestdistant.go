// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildRequestDistant struct{}

func NewGuildRequestDistant(extra string) (GuildRequestDistant, error) {
	var m GuildRequestDistant

	if err := m.Deserialize(extra); err != nil {
		return GuildRequestDistant{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildRequestDistant) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildRequestDistant
}

func (m GuildRequestDistant) MessageName() string {
	return "GuildRequestDistant"
}

func (m GuildRequestDistant) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRequestDistant) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
