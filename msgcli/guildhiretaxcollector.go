// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildHireTaxCollector struct{}

func (m GuildHireTaxCollector) MessageId() retroproto.MsgCliId {
	return retroproto.GuildHireTaxCollector
}

func (m GuildHireTaxCollector) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildHireTaxCollector) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
