// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ExchangeRequestReady struct{}

func (m ExchangeRequestReady) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeRequestReady
}

func (m ExchangeRequestReady) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeRequestReady) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
