// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangePutInInventoryFromShed struct{}

func (m ExchangePutInInventoryFromShed) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangePutInInventoryFromShed
}

func (m ExchangePutInInventoryFromShed) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangePutInInventoryFromShed) Deserialize(extra string) error {
	return nil
}
