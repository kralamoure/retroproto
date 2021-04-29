package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1encoding"
)

type ExchangeBigStoreSearch struct {
	ItemType   d1typ.ItemType
	TemplateId int
}

func (m ExchangeBigStoreSearch) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ExchangeBigStoreSearch
}

func (m ExchangeBigStoreSearch) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.ItemType, m.TemplateId), nil
}

func (m *ExchangeBigStoreSearch) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return d1encoding.ErrInvalidMsg
	}

	itemType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemType = d1typ.ItemType(itemType)

	templateId, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.TemplateId = int(templateId)

	return nil
}
