package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type ExchangeBigStoreItemsList struct {
	TemplateId int
	Items      []typ.ExchangeBigStoreItemsListItem
}

func (m ExchangeBigStoreItemsList) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreItemsList
}

func (m ExchangeBigStoreItemsList) MessageName() string {
	return "ExchangeBigStoreItemsList"
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

	return fmt.Sprintf("%d|%s", m.TemplateId, strings.Join(items, "|")), nil
}

func (m *ExchangeBigStoreItemsList) Deserialize(extra string) error {
	if extra == "" {
		return retroproto.ErrInvalidMsg
	}

	sli := strings.Split(extra, "|")

	itemTemplateId, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.TemplateId = int(itemTemplateId)

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
