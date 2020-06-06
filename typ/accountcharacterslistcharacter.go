package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type AccountCharactersListCharacter struct {
	Id          int
	Name        string
	Level       int
	GFXId       int
	Color1      string
	Color2      string
	Color3      string
	Accessories CommonAccessories
	Merchant    bool
	ServerId    int
	Dead        bool
	DeathCount  int
	LvlMax      int
}

func (m AccountCharactersListCharacter) Serialized() (string, error) {
	var merchant int
	if m.Merchant {
		merchant = 1
	}

	var dead int
	if m.Dead {
		dead = 1
	}

	accessories, err := m.Accessories.Serialized()
	if err != nil {
		return "", err
	}

	color1 := m.Color1
	if color1 == "" {
		color1 = "-1"
	}
	color2 := m.Color2
	if color2 == "" {
		color2 = "-1"
	}
	color3 := m.Color3
	if color3 == "" {
		color3 = "-1"
	}

	return fmt.Sprintf("%d;%s;%d;%d;%s;%s;%s;%s;%d;%d;%d;%d;%d",
		m.Id, m.Name, m.Level, m.GFXId, color1, color2, color3, accessories, merchant, m.ServerId, dead,
		m.DeathCount, m.LvlMax), nil
}

func (m *AccountCharactersListCharacter) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) < 13 {
		return d1proto.ErrInvalidMsg
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
		gfxId, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.GFXId = int(gfxId)
	}

	m.Color1 = sli[4]
	m.Color2 = sli[5]
	m.Color3 = sli[6]

	err := m.Accessories.Deserialize(sli[7])
	if err != nil {
		return err
	}

	if sli[8] != "" {
		merchant, err := strconv.ParseBool(sli[8])
		if err != nil {
			return err
		}
		m.Merchant = merchant
	}

	if sli[9] != "" {
		serverId, err := strconv.ParseInt(sli[9], 10, 32)
		if err != nil {
			return err
		}
		m.ServerId = int(serverId)
	}

	if sli[10] != "" {
		dead, err := strconv.ParseBool(sli[10])
		if err != nil {
			return err
		}
		m.Dead = dead
	}

	if sli[11] != "" {
		deathCount, err := strconv.ParseInt(sli[11], 10, 32)
		if err != nil {
			return err
		}
		m.DeathCount = int(deathCount)
	}

	if sli[12] != "" {
		lvlMax, err := strconv.ParseInt(sli[12], 10, 32)
		if err != nil {
			return err
		}
		m.LvlMax = int(lvlMax)
	}

	return nil
}
