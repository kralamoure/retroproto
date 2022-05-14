// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosMountPark struct{}

func NewGuildInfosMountPark(extra string) (GuildInfosMountPark, error) {
	var m GuildInfosMountPark

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosMountPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosMountPark) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosMountPark
}

func (m GuildInfosMountPark) MessageName() string {
	return "GuildInfosMountPark"
}

func (m GuildInfosMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
