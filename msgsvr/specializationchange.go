package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type SpecializationChange struct {
	Value int
}

func (m SpecializationChange) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpecializationChange
}

func (m SpecializationChange) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Value), nil
}

func (m *SpecializationChange) Deserialize(extra string) error {
	if len(extra) < 1 {
		return d1proto.ErrInvalidMsg
	}

	value, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Value = int(value)

	return nil
}
