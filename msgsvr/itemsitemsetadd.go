package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ItemsItemSetAdd struct {
	Id                int
	ItemsTemplatesIds []int
	// TODO
	Effects string
}

func (m ItemsItemSetAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsItemSetAdd
}

func (m ItemsItemSetAdd) Serialized() (string, error) {
	itemsTemplatesIds := make([]string, len(m.ItemsTemplatesIds))
	for i, v := range m.ItemsTemplatesIds {
		itemsTemplatesIds[i] = fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%d|%s|%s", m.Id, strings.Join(itemsTemplatesIds, ";"), m.Effects), nil
}

func (m *ItemsItemSetAdd) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return d1proto.ErrInvalidMsg
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if sli[1] != "" {
		sli2 := strings.Split(sli[1], ";")
		m.ItemsTemplatesIds = make([]int, len(sli2))
		for i, v := range sli2 {
			itemTemplateId, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			m.ItemsTemplatesIds[i] = int(itemTemplateId)
		}
	}

	m.Effects = sli[2]

	return nil
}
