// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildGetInfosTaxCollector struct{}

func (m GuildGetInfosTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildGetInfosTaxCollector
}

func (m GuildGetInfosTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildGetInfosTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
