package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type AccountAddCharacter struct {
	Name   string
	Class  int
	Sex    int
	Color1 string
	Color2 string
	Color3 string
}

func (m AccountAddCharacter) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountAddCharacter
}

func (m AccountAddCharacter) Serialized() (string, error) {
	color1 := -1
	if m.Color1 != "" {
		color1N, err := strconv.ParseInt(m.Color1, 16, 32)
		if err != nil {
			return "", err
		}
		color1 = int(color1N)
	}

	color2 := -1
	if m.Color2 != "" {
		color2N, err := strconv.ParseInt(m.Color2, 16, 32)
		if err != nil {
			return "", err
		}
		color2 = int(color2N)
	}

	color3 := -1
	if m.Color3 != "" {
		color3N, err := strconv.ParseInt(m.Color3, 16, 32)
		if err != nil {
			return "", err
		}
		color3 = int(color3N)
	}

	return fmt.Sprintf("|%s|%d|%d|%d|%d|%d",
		m.Name, m.Class, m.Sex, color1, color2, color3,
	), nil
}

func (m *AccountAddCharacter) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 6 {
		return retroproto.ErrInvalidMsg
	}

	m.Name = sli[0]

	class, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Class = int(class)

	sex, err := strconv.ParseInt(sli[2], 10, 32)
	if err != nil {
		return err
	}
	m.Sex = int(sex)

	color1, err := strconv.ParseInt(sli[3], 10, 32)
	if err != nil {
		return err
	}
	if color1 < 0 {
		m.Color1 = ""
	} else {
		m.Color1 = fmt.Sprintf("%06x", color1)
	}

	color2, err := strconv.ParseInt(sli[4], 10, 32)
	if err != nil {
		return err
	}
	if color2 < 0 {
		m.Color2 = ""
	} else {
		m.Color2 = fmt.Sprintf("%06x", color2)
	}

	color3, err := strconv.ParseInt(sli[5], 10, 32)
	if err != nil {
		return err
	}
	if color3 < 0 {
		m.Color3 = ""
	} else {
		m.Color3 = fmt.Sprintf("%06x", color3)
	}

	return nil
}
