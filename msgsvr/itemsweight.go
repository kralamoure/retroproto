package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ItemsWeight struct {
	Current int
	Max     int
}

func (m ItemsWeight) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsWeight
}

func (m ItemsWeight) MessageName() string {
	return "ItemsWeight"
}

func (m ItemsWeight) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Current, m.Max), nil
}

func (m *ItemsWeight) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 2 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		current, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Current = int(current)
	}

	if sli[1] != "" {
		max, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Max = int(max)
	}

	return nil
}
