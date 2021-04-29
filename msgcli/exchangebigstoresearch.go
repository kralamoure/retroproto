package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1proto"
)

type ExchangeBigStoreSearch struct {
	ItemType   d1typ.ItemType
	TemplateId int
}

func (m ExchangeBigStoreSearch) ProtocolId() d1proto.MsgCliId {
	return d1proto.ExchangeBigStoreSearch
}

func (m ExchangeBigStoreSearch) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.ItemType, m.TemplateId), nil
}

func (m *ExchangeBigStoreSearch) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return d1proto.ErrInvalidMsg
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
