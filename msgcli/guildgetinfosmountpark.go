// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildGetInfosMountPark struct{}

func NewGuildGetInfosMountPark(extra string) (GuildGetInfosMountPark, error) {
	var m GuildGetInfosMountPark

	if err := m.Deserialize(extra); err != nil {
		return GuildGetInfosMountPark{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildGetInfosMountPark) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosMountPark
}

func (m GuildGetInfosMountPark) MessageName() string {
	return "GuildGetInfosMountPark"
}

func (m GuildGetInfosMountPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosMountPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
