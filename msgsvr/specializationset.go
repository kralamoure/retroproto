package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type SpecializationSet struct {
	Value int
}

func NewSpecializationSet(extra string) (SpecializationSet, error) {
	var m SpecializationSet

	if err := m.Deserialize(extra); err != nil {
		return SpecializationSet{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m SpecializationSet) MessageId() retroproto.MsgSvrId {
	return retroproto.SpecializationSet
}

func (m SpecializationSet) MessageName() string {
	return "SpecializationSet"
}

func (m SpecializationSet) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Value), nil
}

func (m *SpecializationSet) Deserialize(extra string) error {
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
