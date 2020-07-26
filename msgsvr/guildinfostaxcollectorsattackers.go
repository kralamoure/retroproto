// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GuildInfosTaxCollectorsAttackers struct{}

func (m GuildInfosTaxCollectorsAttackers) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GuildInfosTaxCollectorsAttackers
}

func (m GuildInfosTaxCollectorsAttackers) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsAttackers) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
