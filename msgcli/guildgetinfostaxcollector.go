// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GuildGetInfosTaxCollector struct{}

func (m GuildGetInfosTaxCollector) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GuildGetInfosTaxCollector
}

func (m GuildGetInfosTaxCollector) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GuildGetInfosTaxCollector) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
