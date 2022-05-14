package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type BasicsDate struct {
	Year  int
	Month int
	Day   int
}

func (m BasicsDate) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsDate
}

func (m BasicsDate) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d|%d", m.Year, m.Month-1, m.Day), nil
}

func (m *BasicsDate) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return retroproto.ErrInvalidMsg
	}

	year, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Year = int(year)

	month, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Month = int(month + 1)

	day, err := strconv.ParseInt(sli[2], 10, 32)
	if err != nil {
		return err
	}
	m.Day = int(day)

	return nil
}
