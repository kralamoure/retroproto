package msgsvr

import (
	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1proto"
)

type MountEquipAdd struct {
	Id               int
	ModelId          int
	Ancestors        []int
	Capacities       []int
	Name             string
	Sex              d1typ.Gender
	XP               int
	XPMin            int
	XPMax            int
	Level            int
	Mountable        int
	PodsMax          int
	Wild             bool
	Stamina          int
	StaminaMax       int
	Maturity         int
	MaturityMax      int
	Energy           int
	EnergyMax        int
	Serenity         int
	SerenityMin      int
	SerenityMax      int
	Love             int
	LoveMax          int
	Fecundation      int
	Fecundable       bool
	Effects          []d1typ.Effect
	Tired            int
	TiredMax         int
	Reproductions    int
	ReproduxtionsMax int
}

func (m MountEquipAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipAdd
}

func (m MountEquipAdd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *MountEquipAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
