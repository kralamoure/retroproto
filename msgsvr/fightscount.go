package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type FightsCount struct {
	Value int
}

func NewFightsCount(extra string) (FightsCount, error) {
	var m FightsCount

	if err := m.Deserialize(extra); err != nil {
		return FightsCount{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m FightsCount) MessageId() retroproto.MsgSvrId {
	return retroproto.FightsCount
}

func (m FightsCount) MessageName() string {
	return "FightsCount"
}

func (m FightsCount) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Value), nil
}

func (m *FightsCount) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	value, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Value = int(value)

	return nil
}
