// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreSearch struct{}

func (m ExchangeBigStoreSearch) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreSearch
}

func (m ExchangeBigStoreSearch) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreSearch) Deserialize(extra string) error {
	return nil
}
