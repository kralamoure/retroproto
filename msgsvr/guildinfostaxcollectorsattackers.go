// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosTaxCollectorsAttackers struct{}

func (m GuildInfosTaxCollectorsAttackers) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosTaxCollectorsAttackers
}

func (m GuildInfosTaxCollectorsAttackers) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsAttackers) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
