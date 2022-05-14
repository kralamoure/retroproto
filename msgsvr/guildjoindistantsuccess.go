// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildJoinDistantSuccess struct{}

func NewGuildJoinDistantSuccess(extra string) (GuildJoinDistantSuccess, error) {
	var m GuildJoinDistantSuccess

	if err := m.Deserialize(extra); err != nil {
		return GuildJoinDistantSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildJoinDistantSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildJoinDistantSuccess
}

func (m GuildJoinDistantSuccess) MessageName() string {
	return "GuildJoinDistantSuccess"
}

func (m GuildJoinDistantSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinDistantSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
