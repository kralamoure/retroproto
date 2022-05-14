// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeKillMountInPark struct{}

func (m ExchangeKillMountInPark) MessageId() retroproto.MsgCliId {
	return retroproto.ExchangeKillMountInPark
}

func (m ExchangeKillMountInPark) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeKillMountInPark) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
