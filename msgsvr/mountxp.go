package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type MountXP struct {
	Percent int
}

func NewMountXP(extra string) (MountXP, error) {
	var m MountXP

	if err := m.Deserialize(extra); err != nil {
		return MountXP{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountXP) MessageId() retroproto.MsgSvrId {
	return retroproto.MountXP
}

func (m MountXP) MessageName() string {
	return "MountXP"
}

func (m MountXP) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Percent), nil
}

func (m *MountXP) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	percent, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Percent = int(percent)

	return nil
}
