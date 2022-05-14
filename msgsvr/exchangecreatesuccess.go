package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type ExchangeCreateSuccess struct {
	Type    retrotyp.Exchange
	NPCBuy  ExchangeCreateSuccessNPCBuy
	Paddock ExchangeCreateSuccessPaddock
}

type ExchangeCreateSuccessNPCBuy struct {
	Quantity1     int
	Quantity2     int
	Quantity3     int
	Types         []retrotyp.ItemType
	Fee           float32
	MaxLevel      int
	MaxPerAccount int
	NPCId         int
	MaxHours      int
}

type ExchangeCreateSuccessPaddock struct {
	Shed    []typ.CommonMountData
	Paddock []typ.CommonMountData
}

func (m ExchangeCreateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeCreateSuccess
}

func (m ExchangeCreateSuccess) Serialized() (string, error) {
	var s string
	switch m.Type {
	case retrotyp.ExchangeNPCBuy:
		t := m.NPCBuy

		itemTypes := make([]string, len(t.Types))
		for i, v := range t.Types {
			itemTypes[i] = fmt.Sprintf("%d", v)
		}

		s = fmt.Sprintf("%d,%d,%d;%s;%f;%d;%d;%d;%d",
			t.Quantity1, t.Quantity2, t.Quantity3, strings.Join(itemTypes, ","), t.Fee, t.MaxLevel, t.MaxPerAccount,
			t.NPCId, t.MaxHours)
	case retrotyp.ExchangePaddock: // TODO
		t := m.Paddock

		shed := make([]string, len(t.Shed))
		for i, v := range t.Shed {
			str, err := v.Serialized()
			if err != nil {
				return "", err
			}
			shed[i] = str
		}

		paddock := make([]string, len(t.Paddock))
		for i, v := range t.Paddock {
			str, err := v.Serialized()
			if err != nil {
				return "", err
			}
			paddock[i] = str
		}

		s = fmt.Sprintf("%s~%s", strings.Join(shed, ";"), strings.Join(paddock, ";"))
	default:
		return "", retroproto.ErrNotImplemented
	}

	return fmt.Sprintf("%d|%s", m.Type, s), nil
}

func (m *ExchangeCreateSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	exchangeType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}

	switch retrotyp.Exchange(exchangeType) {
	case retrotyp.ExchangeNPCBuy:
		if len(sli) < 2 {
			return retroproto.ErrInvalidMsg
		}

		var t ExchangeCreateSuccessNPCBuy

		sli := strings.Split(sli[1], ";")
		if len(sli) < 7 {
			return retroproto.ErrInvalidMsg
		}

		sli2 := strings.Split(sli[0], ",")
		if len(sli2) < 3 {
			return retroproto.ErrInvalidMsg
		}

		quantity1, err := strconv.ParseInt(sli2[0], 10, 32)
		if err != nil {
			return err
		}
		t.Quantity1 = int(quantity1)

		quantity2, err := strconv.ParseInt(sli2[1], 10, 32)
		if err != nil {
			return err
		}
		t.Quantity2 = int(quantity2)

		quantity3, err := strconv.ParseInt(sli2[2], 10, 32)
		if err != nil {
			return err
		}
		t.Quantity3 = int(quantity3)

		sli2 = strings.Split(sli[1], ",")
		t.Types = make([]retrotyp.ItemType, len(sli2))
		for i, v := range sli2 {
			itemType, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			t.Types[i] = retrotyp.ItemType(itemType)
		}

		fee, err := strconv.ParseFloat(sli[2], 32)
		if err != nil {
			return err
		}
		t.Fee = float32(fee)

		maxLevel, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		t.MaxLevel = int(maxLevel)

		maxPerAccount, err := strconv.ParseInt(sli[4], 10, 32)
		if err != nil {
			return err
		}
		t.MaxPerAccount = int(maxPerAccount)

		npcId, err := strconv.ParseInt(sli[5], 10, 32)
		if err != nil {
			return err
		}
		t.NPCId = int(npcId)

		maxHours, err := strconv.ParseInt(sli[6], 10, 32)
		if err != nil {
			return err
		}
		t.MaxHours = int(maxHours)

		m.NPCBuy = t
	default:
		return retroproto.ErrNotImplemented
	}

	return nil
}
