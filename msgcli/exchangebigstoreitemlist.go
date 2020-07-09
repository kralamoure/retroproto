package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreItemList struct {
	ItemTemplateId int
}

func (m ExchangeBigStoreItemList) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreItemList
}

func (m ExchangeBigStoreItemList) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.ItemTemplateId), nil
}

func (m *ExchangeBigStoreItemList) Deserialize(extra string) error {
	itemTemplateId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.ItemTemplateId = int(itemTemplateId)
	return nil
}
