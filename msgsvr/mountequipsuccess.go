package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type MountEquipSuccess struct {
	Data typ.CommonMountData
}

func NewMountEquipSuccess(extra string) (MountEquipSuccess, error) {
	var m MountEquipSuccess

	if err := m.Deserialize(extra); err != nil {
		return MountEquipSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountEquipSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.MountEquipSuccess
}

func (m MountEquipSuccess) MessageName() string {
	return "MountEquipSuccess"
}

func (m MountEquipSuccess) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountEquipSuccess) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
