package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type AccountStats struct {
	XP               int
	XPLow            int
	XPHigh           int
	Kama             int
	BonusPoints      int
	BonusPointsSpell int
	Alignment        int
	FakeAlignment    int
	AlignmentLevel   int
	Grade            int
	Honour           int
	Disgrace         int
	AlignmentEnabled bool
	LP               int
	LPMax            int
	Energy           int
	EnergyMax        int
	Initiative       int
	Discernment      int
	Characteristics  map[retrotyp.CharacteristicId]retrotyp.Characteristic
}

func (m AccountStats) ProtocolId() retroproto.MsgSvrId {
	return retroproto.AccountStats
}

func (m AccountStats) Serialized() (string, error) {
	var alignmentEnabled int
	if m.AlignmentEnabled {
		alignmentEnabled = 1
	}

	order := []retrotyp.CharacteristicId{
		retrotyp.CharacteristicIdAP,
		retrotyp.CharacteristicIdMP,
		retrotyp.CharacteristicIdStrength,
		retrotyp.CharacteristicIdVitality,
		retrotyp.CharacteristicIdWisdom,
		retrotyp.CharacteristicIdChance,
		retrotyp.CharacteristicIdAgility,
		retrotyp.CharacteristicIdIntelligence,
		retrotyp.CharacteristicIdRange,
		retrotyp.CharacteristicIdMaxSummonedCreaturesBoost,
		retrotyp.CharacteristicIdDamages,
		retrotyp.CharacteristicIdPhysicalDamages,
		retrotyp.CharacteristicIdWeaponDamagesPercent,
		retrotyp.CharacteristicIdDamagesPercent,
		retrotyp.CharacteristicIdHeals,
		retrotyp.CharacteristicIdTrapDamages,
		retrotyp.CharacteristicIdTrapDamagesPercent,
		retrotyp.CharacteristicIdDamagesReflection,
		retrotyp.CharacteristicIdCriticalHits,
		retrotyp.CharacteristicIdCriticalFailures,
		retrotyp.CharacteristicIdDodgeAP,
		retrotyp.CharacteristicIdDodgeMP,
		retrotyp.CharacteristicIdNeutralResistance,
		retrotyp.CharacteristicIdNeutralResistancePercent,
		retrotyp.CharacteristicIdNeutralResistancePVP,
		retrotyp.CharacteristicIdNeutralResistancePercentPVP,
		retrotyp.CharacteristicIdEarthResistance,
		retrotyp.CharacteristicIdEarthResistancePercent,
		retrotyp.CharacteristicIdEarthResistancePVP,
		retrotyp.CharacteristicIdEarthResistancePercentPVP,
		retrotyp.CharacteristicIdWaterResistance,
		retrotyp.CharacteristicIdWaterResistancePercent,
		retrotyp.CharacteristicIdWaterResistancePVP,
		retrotyp.CharacteristicIdWaterResistancePercentPVP,
		retrotyp.CharacteristicIdAirResistance,
		retrotyp.CharacteristicIdAirResistancePercent,
		retrotyp.CharacteristicIdAirResistancePVP,
		retrotyp.CharacteristicIdAirResistancePercentPVP,
		retrotyp.CharacteristicIdFireResistance,
		retrotyp.CharacteristicIdFireResistancePercent,
		retrotyp.CharacteristicIdFireResistancePVP,
		retrotyp.CharacteristicIdFireResistancePercentPVP,
	}

	stats := make([]string, len(order))
	for i, v := range order {
		v := m.Characteristics[v]
		stats[i] = fmt.Sprintf("%d,%d,%d,%d", v.Base, v.Equipment, v.Feat, v.Boost)
	}

	return fmt.Sprintf("%d,%d,%d|%d|%d|%d|%d~%d,%d,%d,%d,%d,%d|%d,%d|%d,%d|%d|%d|%s",
		m.XP, m.XPLow, m.XPHigh, m.Kama, m.BonusPoints, m.BonusPointsSpell, m.Alignment, m.FakeAlignment,
		m.AlignmentLevel, m.Grade, m.Honour, m.Disgrace, alignmentEnabled, m.LP, m.LPMax, m.Energy,
		m.EnergyMax, m.Initiative, m.Discernment,
		strings.Join(stats, "|"),
	), nil
}

func (m *AccountStats) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 51 {
		return retroproto.ErrInvalidMsg
	}

	sli2 := strings.Split(sli[0], ",")
	if len(sli2) < 3 {
		return retroproto.ErrInvalidMsg
	}

	if sli2[0] != "" {
		xp, err := strconv.ParseInt(sli2[0], 10, 32)
		if err != nil {
			return err
		}
		m.XP = int(xp)
	}

	if sli2[1] != "" {
		xpLow, err := strconv.ParseInt(sli2[1], 10, 32)
		if err != nil {
			return err
		}
		m.XPLow = int(xpLow)
	}

	if sli2[2] != "" {
		xpHigh, err := strconv.ParseInt(sli2[2], 10, 32)
		if err != nil {
			return err
		}
		m.XPHigh = int(xpHigh)
	}

	if sli[1] != "" {
		kama, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Kama = int(kama)
	}

	if sli[2] != "" {
		bonusPoints, err := strconv.ParseInt(sli[2], 10, 32)
		if err != nil {
			return err
		}
		m.BonusPoints = int(bonusPoints)
	}

	if sli[3] != "" {
		bonusPointsSpell, err := strconv.ParseInt(sli[3], 10, 32)
		if err != nil {
			return err
		}
		m.BonusPointsSpell = int(bonusPointsSpell)
	}

	sli3 := strings.Split(sli[4], ",")
	if len(sli3) < 6 {
		return retroproto.ErrInvalidMsg
	}

	if sli3[0] != "" {
		sli4 := strings.Split(sli3[0], "~")
		if len(sli4) < 2 {
			return retroproto.ErrInvalidMsg
		}

		alignment, err := strconv.ParseInt(sli4[0], 10, 32)
		if err != nil {
			return err
		}
		m.Alignment = int(alignment)

		fakeAlignment, err := strconv.ParseInt(sli4[1], 10, 32)
		if err != nil {
			return err
		}
		m.FakeAlignment = int(fakeAlignment)
	}

	if sli3[1] != "" {
		alignmentLevel, err := strconv.ParseInt(sli3[1], 10, 32)
		if err != nil {
			return err
		}
		m.AlignmentLevel = int(alignmentLevel)
	}

	if sli3[2] != "" {
		grade, err := strconv.ParseInt(sli3[2], 10, 32)
		if err != nil {
			return err
		}
		m.Grade = int(grade)
	}

	if sli3[3] != "" {
		honour, err := strconv.ParseInt(sli3[3], 10, 32)
		if err != nil {
			return err
		}
		m.Honour = int(honour)
	}

	if sli3[4] != "" {
		disgrace, err := strconv.ParseInt(sli3[4], 10, 32)
		if err != nil {
			return err
		}
		m.Disgrace = int(disgrace)
	}

	if sli3[5] != "" {
		alignmentEnabled, err := strconv.ParseBool(sli3[5])
		if err != nil {
			return err
		}
		m.AlignmentEnabled = alignmentEnabled
	}

	sli4 := strings.Split(sli[5], ",")
	if len(sli4) < 2 {
		return retroproto.ErrInvalidMsg
	}

	if sli4[0] != "" {
		lp, err := strconv.ParseInt(sli4[0], 10, 32)
		if err != nil {
			return err
		}
		m.LP = int(lp)
	}

	if sli4[1] != "" {
		lpMax, err := strconv.ParseInt(sli4[1], 10, 32)
		if err != nil {
			return err
		}
		m.LPMax = int(lpMax)
	}

	sli5 := strings.Split(sli[6], ",")
	if len(sli5) < 2 {
		return retroproto.ErrInvalidMsg
	}

	if sli5[0] != "" {
		energy, err := strconv.ParseInt(sli5[0], 10, 32)
		if err != nil {
			return err
		}
		m.Energy = int(energy)
	}

	if sli5[1] != "" {
		energyMax, err := strconv.ParseInt(sli5[1], 10, 32)
		if err != nil {
			return err
		}
		m.EnergyMax = int(energyMax)
	}

	if sli[7] != "" {
		initiative, err := strconv.ParseInt(sli[7], 10, 32)
		if err != nil {
			return err
		}
		m.Initiative = int(initiative)
	}

	if sli[8] != "" {
		discernment, err := strconv.ParseInt(sli[8], 10, 32)
		if err != nil {
			return err
		}
		m.Discernment = int(discernment)
	}

	order := []retrotyp.CharacteristicId{
		retrotyp.CharacteristicIdAP,
		retrotyp.CharacteristicIdMP,
		retrotyp.CharacteristicIdStrength,
		retrotyp.CharacteristicIdVitality,
		retrotyp.CharacteristicIdWisdom,
		retrotyp.CharacteristicIdChance,
		retrotyp.CharacteristicIdAgility,
		retrotyp.CharacteristicIdIntelligence,
		retrotyp.CharacteristicIdRange,
		retrotyp.CharacteristicIdMaxSummonedCreaturesBoost,
		retrotyp.CharacteristicIdDamages,
		retrotyp.CharacteristicIdPhysicalDamages,
		retrotyp.CharacteristicIdWeaponDamagesPercent,
		retrotyp.CharacteristicIdDamagesPercent,
		retrotyp.CharacteristicIdHeals,
		retrotyp.CharacteristicIdTrapDamages,
		retrotyp.CharacteristicIdTrapDamagesPercent,
		retrotyp.CharacteristicIdDamagesReflection,
		retrotyp.CharacteristicIdCriticalHits,
		retrotyp.CharacteristicIdCriticalFailures,
		retrotyp.CharacteristicIdDodgeAP,
		retrotyp.CharacteristicIdDodgeMP,
		retrotyp.CharacteristicIdNeutralResistance,
		retrotyp.CharacteristicIdNeutralResistancePercent,
		retrotyp.CharacteristicIdNeutralResistancePVP,
		retrotyp.CharacteristicIdNeutralResistancePercentPVP,
		retrotyp.CharacteristicIdEarthResistance,
		retrotyp.CharacteristicIdEarthResistancePercent,
		retrotyp.CharacteristicIdEarthResistancePVP,
		retrotyp.CharacteristicIdEarthResistancePercentPVP,
		retrotyp.CharacteristicIdWaterResistance,
		retrotyp.CharacteristicIdWaterResistancePercent,
		retrotyp.CharacteristicIdWaterResistancePVP,
		retrotyp.CharacteristicIdWaterResistancePercentPVP,
		retrotyp.CharacteristicIdAirResistance,
		retrotyp.CharacteristicIdAirResistancePercent,
		retrotyp.CharacteristicIdAirResistancePVP,
		retrotyp.CharacteristicIdAirResistancePercentPVP,
		retrotyp.CharacteristicIdFireResistance,
		retrotyp.CharacteristicIdFireResistancePercent,
		retrotyp.CharacteristicIdFireResistancePVP,
		retrotyp.CharacteristicIdFireResistancePercentPVP,
	}

	m.Characteristics = make(map[retrotyp.CharacteristicId]retrotyp.Characteristic, len(order))
	for i, v := range order {
		sli2 := strings.Split(sli[i+9], ",")
		if len(sli2) < 4 {
			return retroproto.ErrInvalidMsg
		}

		base, err := strconv.ParseInt(sli2[0], 10, 32)
		if err != nil {
			return err
		}
		equipment, err := strconv.ParseInt(sli2[1], 10, 32)
		if err != nil {
			return err
		}
		feat, err := strconv.ParseInt(sli2[2], 10, 32)
		if err != nil {
			return err
		}
		boost, err := strconv.ParseInt(sli2[3], 10, 32)
		if err != nil {
			return err
		}

		m.Characteristics[v] = retrotyp.Characteristic{
			Id:        v,
			Base:      int(base),
			Equipment: int(equipment),
			Feat:      int(feat),
			Boost:     int(boost),
		}
	}

	return nil
}
