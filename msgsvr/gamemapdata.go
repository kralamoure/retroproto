package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type GameMapData struct {
	Id   int
	Name string
	Key  string
}

func NewGameMapData(extra string) (GameMapData, error) {
	var m GameMapData

	if err := m.Deserialize(extra); err != nil {
		return GameMapData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameMapData) MessageId() retroproto.MsgSvrId {
	return retroproto.GameMapData
}

func (m GameMapData) MessageName() string {
	return "GameMapData"
}

func (m GameMapData) Serialized() (string, error) {
	return fmt.Sprintf("|%d|%s|%s", m.Id, m.Name, m.Key), nil
}

func (m *GameMapData) Deserialize(extra string) error {
	extra = strings.TrimPrefix(extra, "|")

	sli := strings.Split(extra, "|")
	if len(sli) < 3 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	m.Name = sli[1]
	m.Key = sli[2]

	return nil
}
