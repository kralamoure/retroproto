// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildNew struct{}

func NewGuildNew(extra string) (GuildNew, error) {
	var m GuildNew

	if err := m.Deserialize(extra); err != nil {
		return GuildNew{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildNew) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildNew
}

func (m GuildNew) MessageName() string {
	return "GuildNew"
}

func (m GuildNew) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildNew) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
