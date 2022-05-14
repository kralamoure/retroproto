package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreTypeItemsList struct {
	ItemType        retrotyp.ItemType
	ItemTemplateIds []int
}

func (m ExchangeBigStoreTypeItemsList) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeBigStoreTypeItemsList
}

func (m ExchangeBigStoreTypeItemsList) MessageName() string {
	return "ExchangeBigStoreTypeItemsList"
}

func (m ExchangeBigStoreTypeItemsList) Serialized() (string, error) {
	itemTemplateIds := make([]string, len(m.ItemTemplateIds))
	for i, v := range m.ItemTemplateIds {
		itemTemplateIds[i] = fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%d|%s", m.ItemType, strings.Join(itemTemplateIds, ";")), nil
}

func (m *ExchangeBigStoreTypeItemsList) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
	}

	itemType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemType = retrotyp.ItemType(itemType)

	if sli[1] != "" {
		itemTemplateIds := strings.Split(sli[1], ";")
		m.ItemTemplateIds = make([]int, len(itemTemplateIds))
		for i, v := range itemTemplateIds {
			itemTemplateId, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			m.ItemTemplateIds[i] = int(itemTemplateId)
		}
	}

	return nil
}
