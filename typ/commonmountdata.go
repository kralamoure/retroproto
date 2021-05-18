package typ

import (
	"fmt"
	"strings"

	"github.com/kralamoure/retro"
	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type CommonMountData struct {
	Id               int
	ModelId          int
	Ancestors        [14]int
	Capacities       []retrotyp.MountCapacityId
	Name             string
	Sex              retrotyp.Sex
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
	Effects          []retrotyp.Effect
	Tiredness        int
	TirednessMax     int
	Reproductions    int
	ReproductionsMax int
}

func (m CommonMountData) Serialized() (string, error) {
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

	effects := retro.EncodeItemEffects(m.Effects)

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
		strings.Join(effects, ","),
		m.Tiredness,
		m.TirednessMax,
		m.Reproductions,
		m.ReproductionsMax,
	), nil
}

func (m *CommonMountData) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
