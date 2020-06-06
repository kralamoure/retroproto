// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreBuy struct{}

func (m ExchangeBigStoreBuy) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreBuy
}

func (m ExchangeBigStoreBuy) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreBuy) Deserialize(extra string) error {
	return nil
}
