package msgsvr

import (
	"fmt"
	"strings"

	"github.com/kralamoure/d1"
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
	Mountable        bool
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
	Tiredness        int
	TirednessMax     int
	Reproductions    int
	ReproductionsMax int
}

func (m MountEquipAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountEquipAdd
}

func (m MountEquipAdd) Serialized() (string, error) {
	ancestors := make([]string, len(m.Ancestors))
	for i := range m.Ancestors {
		ancestors[i] = fmt.Sprintf("%d", m.Ancestors[i])
	}

	capacities := make([]string, len(m.Capacities))
	for i := range m.Capacities {
		capacities[i] = fmt.Sprintf("%d", m.Capacities[i])
	}

	var mountable int
	if m.Mountable {
		mountable = 1
	}

	var wild int
	if m.Wild {
		wild = 1
	}

	var fecundable int
	if m.Fecundable {
		fecundable = 1
	}

	effects := d1.EncodeItemEffects(m.Effects)

	return fmt.Sprintf("%d:%d:%s:%s:%s:%d:%d,%d,%d:%d:%d:%d:%d:%d,%d:%d,%d:%d,%d:%d,%d,%d:%d,%d:%d:%d:%s:%d,%d:%d,%d",
		m.Id,
		m.ModelId,
		strings.Join(ancestors, ","),
		strings.Join(capacities, ","),
		m.Name,
		m.Sex,
		m.XP,
		m.XPMin,
		m.XPMax,
		m.Level,
		mountable,
		m.PodsMax,
		wild,
		m.Stamina,
		m.StaminaMax,
		m.Maturity,
		m.MaturityMax,
		m.Energy,
		m.EnergyMax,
		m.Serenity,
		m.SerenityMin,
		m.SerenityMax,
		m.Love,
		m.LoveMax,
		m.Fecundation,
		fecundable,
		effects,
		m.Tiredness,
		m.TirednessMax,
		m.Reproductions,
		m.ReproductionsMax,
	), nil
}

func (m *MountEquipAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
