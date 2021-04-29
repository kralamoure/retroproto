package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1"
	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1encoding"
)

type ItemsItemSetAdd struct {
	Id                int
	ItemsTemplatesIds []int
	Effects           []d1typ.Effect
}

func (m ItemsItemSetAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ItemsItemSetAdd
}

func (m ItemsItemSetAdd) Serialized() (string, error) {
	itemsTemplatesIds := make([]string, len(m.ItemsTemplatesIds))
	for i, v := range m.ItemsTemplatesIds {
		itemsTemplatesIds[i] = fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%d|%s|%s", m.Id, strings.Join(itemsTemplatesIds, ";"), strings.Join(d1.EncodeItemEffects(m.Effects), ",")), nil
}

func (m *ItemsItemSetAdd) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return d1encoding.ErrInvalidMsg
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

	effectsStr := sli[2]
	effects, err := d1.DecodeItemEffects(strings.Split(effectsStr, ","))
	if err != nil {
		return err
	}
	m.Effects = effects

	return nil
}
