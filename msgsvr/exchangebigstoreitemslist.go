package msgsvr

import (
	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/typ"
)

type ExchangeBigStoreItemsList struct {
	ItemTemplateId int
	Items          []typ.ExchangeBigStoreItemsListItem
}

func (m ExchangeBigStoreItemsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ExchangeBigStoreItemsList
}

func (m ExchangeBigStoreItemsList) Serialized() (string, error) {
	return "", nil
}

func (m *ExchangeBigStoreItemsList) Deserialize(extra string) error {
	return nil
}
