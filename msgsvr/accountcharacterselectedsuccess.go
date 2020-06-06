package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type AccountCharacterSelectedSuccess struct {
	Id     int
	Name   string
	Level  int
	Guild  int
	Sex    int
	GFXId  int
	Color1 string
	Color2 string
	Color3 string
	// TODO
	Items string
}

func (m AccountCharacterSelectedSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountCharacterSelectedSuccess
}

func (m AccountCharacterSelectedSuccess) Serialized() (string, error) {
	return fmt.Sprintf("|%d|%s|%d|%d|%d|%d|%s|%s|%s|%s",
		m.Id, m.Name, m.Level, m.Guild, m.Sex, m.GFXId, m.Color1, m.Color2, m.Color3, m.Items), nil
}

func (m *AccountCharacterSelectedSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 10 {
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
		guild, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.Guild = int(guild)
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
	m.Items = sli[9]

	return nil
}
