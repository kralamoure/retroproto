package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type CommonGuildEmblem struct {
	BackId    int
	BackColor string
	UpId      int
	UpColor   string
}

func NewCommonGuildEmblem(extra string) (CommonGuildEmblem, error) {
	var m CommonGuildEmblem

	if err := m.Deserialize(extra); err != nil {
		return CommonGuildEmblem{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m CommonGuildEmblem) Serialized() (string, error) {
	if m.BackId == 0 && m.UpId == 0 && m.BackColor == "" && m.UpColor == "" {
		return "", nil
	}

	backId := strconv.FormatInt(int64(m.BackId), 36)
	upId := strconv.FormatInt(int64(m.UpId), 36)

	backColor := ""
	if m.BackColor != "" {
		backColorN, err := strconv.ParseInt(m.BackColor, 16, 32)
		if err != nil {
			return "", err
		}
		backColor = strconv.FormatInt(backColorN, 36)
	}

	upColor := ""
	if m.UpColor != "" {
		upColorN, err := strconv.ParseInt(m.UpColor, 16, 32)
		if err != nil {
			return "", err
		}
		upColor = strconv.FormatInt(upColorN, 36)
	}

	accessories := fmt.Sprintf("%s,%s,%s,%s",
		backId, backColor, upId, upColor,
	)

	return accessories, nil
}

func (m *CommonGuildEmblem) Deserialize(extra string) error {
	if extra == "" {
		return nil
	}

	sli := strings.Split(extra, ",")
	if len(sli) < 4 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		backId, err := strconv.ParseInt(sli[0], 36, 32)
		if err != nil {
			return err
		}
		m.BackId = int(backId)
	}

	if sli[2] != "" {
		upId, err := strconv.ParseInt(sli[2], 36, 32)
		if err != nil {
			return err
		}
		m.UpId = int(upId)
	}

	if sli[1] != "" {
		backColorN, err := strconv.ParseInt(sli[1], 36, 32)
		if err != nil {
			return err
		}

		if backColorN > 0 {
			m.BackColor = strconv.FormatInt(backColorN, 16)
		}
	}

	if sli[3] != "" {
		upColorN, err := strconv.ParseInt(sli[3], 36, 32)
		if err != nil {
			return err
		}

		if upColorN > 0 {
			m.UpColor = strconv.FormatInt(upColorN, 16)
		}
	}

	return nil
}
