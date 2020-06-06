package msgsvr

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kralamoure/d1proto"
)

type GameMapData struct {
	Id   int
	Date time.Time
	Key  string
}

func (m GameMapData) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GameMapData
}

func (m GameMapData) Serialized() (string, error) {
	return fmt.Sprintf("|%d|%s|%s", m.Id, m.Date.UTC().Format("0601021504"), m.Key), nil
}

func (m *GameMapData) Deserialize(extra string) error {
	extra = strings.TrimPrefix(extra, "|")

	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return d1proto.ErrInvalidMsg
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if sli[1] != "" {
		date, err := time.Parse("0601021504", sli[1])
		if err != nil {
			return err
		}
		m.Date = date
	}

	m.Key = sli[1]

	return nil
}
