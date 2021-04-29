package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/typ"
)

type ExchangeCreateSuccess struct {
	Type    d1typ.Exchange
	NPCBuy  ExchangeCreateSuccessNPCBuy
	Paddock ExchangeCreateSuccessPaddock
}

type ExchangeCreateSuccessNPCBuy struct {
	Quantity1     int
	Quantity2     int
	Quantity3     int
	Types         []d1typ.ItemType
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

func (m ExchangeCreateSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ExchangeCreateSuccess
}

func (m ExchangeCreateSuccess) Serialized() (string, error) {
	var s string
	switch m.Type {
	case d1typ.ExchangeNPCBuy:
		t := m.NPCBuy

		itemTypes := make([]string, len(t.Types))
		for i, v := range t.Types {
			itemTypes[i] = fmt.Sprintf("%d", v)
		}

		s = fmt.Sprintf("%d,%d,%d;%s;%f;%d;%d;%d;%d",
			t.Quantity1, t.Quantity2, t.Quantity3, strings.Join(itemTypes, ","), t.Fee, t.MaxLevel, t.MaxPerAccount,
			t.NPCId, t.MaxHours)
	case d1typ.ExchangePaddock: // TODO
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
		return "", d1encoding.ErrNotImplemented
	}

	return fmt.Sprintf("%d|%s", m.Type, s), nil
}

func (m *ExchangeCreateSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	exchangeType, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}

	switch d1typ.Exchange(exchangeType) {
	case d1typ.ExchangeNPCBuy:
		if len(sli) < 2 {
			return d1encoding.ErrInvalidMsg
		}

		var t ExchangeCreateSuccessNPCBuy

		sli := strings.Split(sli[1], ";")
		if len(sli) < 7 {
			return d1encoding.ErrInvalidMsg
		}

		sli2 := strings.Split(sli[0], ",")
		if len(sli2) < 3 {
			return d1encoding.ErrInvalidMsg
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
		t.Types = make([]d1typ.ItemType, len(sli2))
		for i, v := range sli2 {
			itemType, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return err
			}
			t.Types[i] = d1typ.ItemType(itemType)
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
		return d1encoding.ErrNotImplemented
	}

	return nil
}
