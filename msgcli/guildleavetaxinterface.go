// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GuildLeaveTaxInterface struct{}

func (m GuildLeaveTaxInterface) MessageId() retroproto.MsgCliId {
	return retroproto.GuildLeaveTaxInterface
}

func (m GuildLeaveTaxInterface) MessageName() string {
	return "GuildLeaveTaxInterface"
}

func (m GuildLeaveTaxInterface) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GuildLeaveTaxInterface) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
