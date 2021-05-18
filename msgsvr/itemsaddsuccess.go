package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro"
	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/enum"
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
	Position   retrotyp.CharacterItemPosition
	Effects    []retrotyp.Effect
}

func (m ItemsAddSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ItemsAddSuccess
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
				effects := retro.EncodeItemEffects(v.Effects)
				objects[i] = fmt.Sprintf("%x~%x~%x~%x~%s", v.Id, v.TemplateId, v.Quantity, int(v.Position), strings.Join(effects, ","))
			}
			items[i] = fmt.Sprintf("%s%s", string(v.ItemType), strings.Join(objects, ";"))
		default:
			return "", retroproto.ErrInvalidMsg
		}
	}
	return strings.Join(items, "*"), nil
}

func (m *ItemsAddSuccess) Deserialize(extra string) error {
	items := strings.Split(extra, "*")
	m.Items = make([]ItemsAddSuccessItem, len(items))
	for i, v := range items {
		if len([]rune(v)) < 2 {
			return retroproto.ErrInvalidMsg
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
					return retroproto.ErrInvalidMsg
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

				quantity, err := strconv.ParseInt(sli[2], 16, 32)
				if err != nil {
					return err
				}
				object.Quantity = int(quantity)

				if sli[3] == "" {
					object.Position = -1
				} else {
					position, err := strconv.ParseInt(sli[3], 16, 32)
					if err != nil {
						return err
					}
					object.Position = retrotyp.CharacterItemPosition(position)
				}
				effects, err := retro.DecodeItemEffects(strings.Split(sli[4], ","))
				if err != nil {
					return err
				}
				object.Effects = effects

				item.Objects[i] = object
			}
		default:
			return retroproto.ErrInvalidMsg
		}
		m.Items[i] = item
	}

	return nil
}
