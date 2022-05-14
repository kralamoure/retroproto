// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GuildInfosTaxCollectorsAttackers struct{}

func (m GuildInfosTaxCollectorsAttackers) MessageId() retroproto.MsgSvrId {
	return retroproto.GuildInfosTaxCollectorsAttackers
}

func (m GuildInfosTaxCollectorsAttackers) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsAttackers) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
