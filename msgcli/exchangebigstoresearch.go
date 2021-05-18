package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type ExchangeBigStoreSearch struct {
	ItemType   retrotyp.ItemType
	TemplateId int
}

func (m ExchangeBigStoreSearch) ProtocolId() retroproto.MsgCliId {
	return retroproto.ExchangeBigStoreSearch
}

func (m ExchangeBigStoreSearch) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.ItemType, m.TemplateId), nil
}

func (m *ExchangeBigStoreSearch) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	itemType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.ItemType = retrotyp.ItemType(itemType)

	templateId, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.TemplateId = int(templateId)

	return nil
}
