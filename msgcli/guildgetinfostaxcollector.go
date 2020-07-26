// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GuildGetInfosTaxCollector struct{}

func (m GuildGetInfosTaxCollector) ProtocolId() d1proto.MsgCliId {
	return d1proto.GuildGetInfosTaxCollector
}

func (m GuildGetInfosTaxCollector) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GuildGetInfosTaxCollector) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
