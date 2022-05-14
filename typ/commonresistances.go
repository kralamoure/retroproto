package typ

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type CommonResistances struct {
	ResistanceNeutral int
	ResistanceEarth   int
	ResistanceWater   int
	ResistanceAir     int
	ResistanceFire    int
	DodgeAP           int
	DodgeMP           int
}

func NewCommonResistances(extra string) (CommonResistances, error) {
	var m CommonResistances

	if err := m.Deserialize(extra); err != nil {
		return CommonResistances{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m CommonResistances) Serialized() (string, error) {
	resistances := fmt.Sprintf("%d;%d;%d;%d;%d;%d;%d",
		m.ResistanceNeutral, m.ResistanceEarth, m.ResistanceWater, m.ResistanceAir, m.ResistanceFire,
		m.DodgeAP, m.DodgeMP,
	)

	return resistances, nil
}

func (m *CommonResistances) Deserialize(extra string) error {
	sli := strings.Split(extra, ";")
	if len(sli) < 7 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "" {
		resistanceNeutral, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.ResistanceNeutral = int(resistanceNeutral)
	}

	if sli[1] != "" {
		resistanceEarth, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.ResistanceEarth = int(resistanceEarth)
	}

	if sli[2] != "" {
		resistanceWater, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.ResistanceWater = int(resistanceWater)
	}

	if sli[3] != "" {
		resistanceAir, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.ResistanceAir = int(resistanceAir)
	}

	if sli[4] != "" {
		resistanceFire, err := strconv.ParseInt(sli[4], 10, 32)
		if err != nil {
			return err
		}
		m.ResistanceFire = int(resistanceFire)
	}

	if sli[5] != "" {
		dodgeAP, err := strconv.ParseInt(sli[5], 10, 32)
		if err != nil {
			return err
		}
		m.DodgeAP = int(dodgeAP)
	}

	if sli[6] != "" {
		dodgeMP, err := strconv.ParseInt(sli[6], 10, 32)
		if err != nil {
			return err
		}
		m.DodgeMP = int(dodgeMP)
	}

	return nil
}
