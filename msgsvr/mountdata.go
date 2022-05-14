package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/typ"
)

type MountData struct {
	Data typ.CommonMountData
}

func NewMountData(extra string) (MountData, error) {
	var m MountData

	if err := m.Deserialize(extra); err != nil {
		return MountData{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m MountData) MessageId() retroproto.MsgSvrId {
	return retroproto.MountData
}

func (m MountData) MessageName() string {
	return "MountData"
}

func (m MountData) Serialized() (string, error) {
	return m.Data.Serialized()
}

func (m *MountData) Deserialize(extra string) error {
	return m.Data.Deserialize(extra)
}
