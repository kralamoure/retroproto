// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildInfosGeneral struct{}

func NewGuildInfosGeneral(extra string) (GuildInfosGeneral, error) {
	var m GuildInfosGeneral

	if err := m.Deserialize(extra); err != nil {
		return GuildInfosGeneral{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildInfosGeneral) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosGeneral
}

func (m GuildInfosGeneral) MessageName() string {
	return "GuildInfosGeneral"
}

func (m GuildInfosGeneral) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosGeneral) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
