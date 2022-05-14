// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildRequestLocal struct{}

func NewGuildRequestLocal(extra string) (GuildRequestLocal, error) {
	var m GuildRequestLocal

	if err := m.Deserialize(extra); err != nil {
		return GuildRequestLocal{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildRequestLocal) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildRequestLocal
}

func (m GuildRequestLocal) MessageName() string {
	return "GuildRequestLocal"
}

func (m GuildRequestLocal) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildRequestLocal) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
