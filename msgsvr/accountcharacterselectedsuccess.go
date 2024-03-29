package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type AccountCharacterSelectedSuccess struct {
	Id      int
	Name    string
	Level   int
	ClassId retrotyp.ClassId // "guild" in client sources, but it's actually class
	Sex     int
	GFXId   int
	Color1  string
	Color2  string
	Color3  string
	Items   []typ.AccountCharacterSelectedSuccessItem
}

func NewAccountCharacterSelectedSuccess(extra string) (AccountCharacterSelectedSuccess, error) {
	var m AccountCharacterSelectedSuccess

	if err := m.Deserialize(extra); err != nil {
		return AccountCharacterSelectedSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountCharacterSelectedSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.AccountCharacterSelectedSuccess
}

func (m AccountCharacterSelectedSuccess) MessageName() string {
	return "AccountCharacterSelectedSuccess"
}

func (m AccountCharacterSelectedSuccess) Serialized() (string, error) {
	items := make([]string, len(m.Items))
	for i, v := range m.Items {
		item, err := v.Serialized()
		if err != nil {
			return "", err
		}
		items[i] = item
	}

	return fmt.Sprintf("|%d|%s|%d|%d|%d|%d|%s|%s|%s|%s",
		m.Id, m.Name, m.Level, m.ClassId, m.Sex, m.GFXId, m.Color1, m.Color2, m.Color3, strings.Join(items, ";")), nil
}

func (m *AccountCharacterSelectedSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 10 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	m.Name = sli[1]

	if sli[2] != "" {
		level, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.Level = int(level)
	}

	if sli[3] != "" {
		classId, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.ClassId = retrotyp.ClassId(classId)
	}

	if sli[4] != "" {
		sex, err := strconv.ParseInt(sli[4], 10, 32)
		if err != nil {
			return err
		}
		m.Sex = int(sex)
	}

	if sli[5] != "" {
		gfxId, err := strconv.ParseInt(sli[5], 10, 32)
		if err != nil {
			return err
		}
		m.GFXId = int(gfxId)
	}

	m.Color1 = sli[6]
	m.Color2 = sli[7]
	m.Color3 = sli[8]

	if sli[9] != "" {
		items := strings.Split(sli[9], ";")
		m.Items = make([]typ.AccountCharacterSelectedSuccessItem, len(items))
		for i, v := range items {
			var item typ.AccountCharacterSelectedSuccessItem
			err := item.Deserialize(v)
			if err != nil {
				return err
			}
			m.Items[i] = item
		}
	}

	return nil
}
