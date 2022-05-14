// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GuildJoinSuccess struct{}

func NewGuildJoinSuccess(extra string) (GuildJoinSuccess, error) {
	var m GuildJoinSuccess

	if err := m.Deserialize(extra); err != nil {
		return GuildJoinSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GuildJoinSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildJoinSuccess
}

func (m GuildJoinSuccess) MessageName() string {
	return "GuildJoinSuccess"
}

func (m GuildJoinSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildJoinSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
