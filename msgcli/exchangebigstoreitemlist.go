// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemList struct{}

func (m ExchangeBigStoreItemList) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreItemList
}

func (m ExchangeBigStoreItemList) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreItemList) Deserialize(extra string) error {
	return nil
}
