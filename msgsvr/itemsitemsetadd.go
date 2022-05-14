package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro"
	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type ItemsItemSetAdd struct {
	Id                int
	ItemsTemplatesIds []int
	Effects           []retrotyp.Effect
}

func NewItemsItemSetAdd(extra string) (ItemsItemSetAdd, error) {
	var m ItemsItemSetAdd

	if err := m.Deserialize(extra); err != nil {
		return ItemsItemSetAdd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsItemSetAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsItemSetAdd
}

func (m ItemsItemSetAdd) MessageName() string {
	return "ItemsItemSetAdd"
}

func (m ItemsItemSetAdd) Serialized() (string, error) {
	itemsTemplatesIds := make([]string, len(m.ItemsTemplatesIds))
	for i, v := range m.ItemsTemplatesIds {
		itemsTemplatesIds[i] = fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%d|%s|%s", m.Id, strings.Join(itemsTemplatesIds, ";"), strings.Join(retro.EncodeItemEffects(m.Effects), ",")), nil
}

func (m *ItemsItemSetAdd) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return retroproto.ErrInvalidMsg
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
	effects, err := retro.DecodeItemEffects(strings.Split(effectsStr, ","))
	if err != nil {
		return err
	}
	m.Effects = effects

	return nil
}
