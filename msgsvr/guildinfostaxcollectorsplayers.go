// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GuildInfosTaxCollectorsPlayers struct{}

func (m GuildInfosTaxCollectorsPlayers) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GuildInfosTaxCollectorsPlayers
}

func (m GuildInfosTaxCollectorsPlayers) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildInfosTaxCollectorsPlayers) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
