package msgsvr

import (
	"fmt"
	"strings"
	"time"

	"github.com/kralamoure/d1proto"
)

type BasicsDate struct {
	Value time.Time
}

func (m BasicsDate) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsDate
}

func (m BasicsDate) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d|%d", m.Value.Year(), m.Value.Month()-1, m.Value.Day()), nil
}

func (m *BasicsDate) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return d1proto.ErrInvalidMsg
	}

	value, err := time.Parse("2006|1|2", fmt.Sprintf("%s|%s|%s", sli[0], sli[1], sli[2]))
	if err != nil {
		return err
	}
	m.Value = value

	return nil
}
