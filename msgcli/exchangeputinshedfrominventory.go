// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePutInShedFromInventory struct{}

func (m ExchangePutInShedFromInventory) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInShedFromInventory
}

func (m ExchangePutInShedFromInventory) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ExchangePutInShedFromInventory) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
