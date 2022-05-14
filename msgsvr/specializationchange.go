package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type SpecializationChange struct {
	Value int
}

func NewSpecializationChange(extra string) (SpecializationChange, error) {
	var m SpecializationChange

	if err := m.Deserialize(extra); err != nil {
		return SpecializationChange{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpecializationChange) MessageId() retroproto.MsgSvrId {
	return retroproto.SpecializationChange
}

func (m SpecializationChange) MessageName() string {
	return "SpecializationChange"
}

func (m SpecializationChange) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Value), nil
}

func (m *SpecializationChange) Deserialize(extra string) error {
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
