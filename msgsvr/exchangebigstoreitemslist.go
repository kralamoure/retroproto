package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

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
	items := make([]string, len(m.Items))
	for i, v := range m.Items {
		item, err := v.Serialized()
		if err != nil {
			return "", err
		}
		items[i] = item
	}

	return fmt.Sprintf("%d|%s", m.ItemTemplateId, strings.Join(items, "|")), nil
}

func (m *ExchangeBigStoreItemsList) Deserialize(extra string) error {
	if extra == "" {
		return d1proto.ErrInvalidMsg
	}

	sli := strings.Split(extra, "|")

	itemTemplateId, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemTemplateId = int(itemTemplateId)

	if len(sli) > 1 {
		m.Items = make([]typ.ExchangeBigStoreItemsListItem, len(sli[1:]))
		for i, v := range sli[1:] {
			err := m.Items[i].Deserialize(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
