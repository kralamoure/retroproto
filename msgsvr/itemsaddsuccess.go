package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1"
	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/enum"
)

type ItemsAddSuccess struct {
	Items []ItemsAddSuccessItem
}

type ItemsAddSuccessItem struct {
	// Either 'G' (I don't know what G stands for) or 'O' (Object).
	ItemType rune
	Objects  []ItemsAddSuccessItemObject
}

type ItemsAddSuccessItemObject struct {
	Id         int
	TemplateId int
	Quantity   int
	Position   int
	Effects    []d1typ.Effect
}

func (m ItemsAddSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsAddSuccess
}

func (m ItemsAddSuccess) Serialized() (string, error) {
	items := make([]string, len(m.Items))
	for i, v := range m.Items {
		switch v.ItemType {
		case enum.ItemsAddSuccessItemType.G:
			items[i] = string(enum.ItemsAddSuccessItemType.G)
		case enum.ItemsAddSuccessItemType.Objects:
			objects := make([]string, len(v.Objects))
			for i, v := range v.Objects {
				effects := d1.EncodeEffects(v.Effects)
				objects[i] = fmt.Sprintf("%x~%x~%x~%x~%s", v.Id, v.TemplateId, v.Quantity, v.Position, effects)
			}
			items[i] = fmt.Sprintf("%s%s", string(v.ItemType), strings.Join(objects, ";"))
		default:
			return "", d1proto.ErrInvalidMsg
		}
	}
	return strings.Join(items, "*"), nil
}

func (m *ItemsAddSuccess) Deserialize(extra string) error {
	items := strings.Split(extra, "*")
	m.Items = make([]ItemsAddSuccessItem, len(items))
	for i, v := range items {
		if len([]rune(v)) < 2 {
			return d1proto.ErrInvalidMsg
		}

		item := ItemsAddSuccessItem{
			ItemType: []rune(v)[0],
		}

		switch item.ItemType {
		case enum.ItemsAddSuccessItemType.G:
		case enum.ItemsAddSuccessItemType.Objects:
			objects := strings.Split(string([]rune(v)[1:]), ";")
			item.Objects = make([]ItemsAddSuccessItemObject, len(objects))
			for i, v := range objects {
				sli := strings.Split(v, "~")
				if len(sli) < 5 {
					return d1proto.ErrInvalidMsg
				}

				var object ItemsAddSuccessItemObject

				id, err := strconv.ParseInt(sli[0], 16, 32)
				if err != nil {
					return err
				}
				object.Id = int(id)

				templateId, err := strconv.ParseInt(sli[1], 16, 32)
				if err != nil {
					return err
				}
				object.TemplateId = int(templateId)

				quantity, err := strconv.ParseInt(sli[3], 16, 32)
				if err != nil {
					return err
				}
				object.Quantity = int(quantity)

				if sli[4] == "" {
					object.Position = -1
				} else {
					position, err := strconv.ParseInt(sli[4], 16, 32)
					if err != nil {
						return err
					}
					object.Position = int(position)
				}
				effects, err := d1.DecodeEffects(sli[5])
				if err != nil {
					return err
				}
				object.Effects = effects

				item.Objects[i] = object
			}
		default:
			return d1proto.ErrInvalidMsg
		}
		m.Items[i] = item
	}

	return nil
}
