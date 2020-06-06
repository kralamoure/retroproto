package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/typ"
)

type AccountStats struct {
	XP                          int
	XPLow                       int
	XPHigh                      int
	Kama                        int
	BonusPoints                 int
	BonusPointsSpell            int
	Alignment                   int
	FakeAlignment               int
	AlignmentUnknownValue       int // TODO: maybe check dofus/utils/DofusTranslator.as
	RankValue                   int
	Honour                      int
	Disgrace                    int
	AlignmentEnabled            bool
	LP                          int
	LPMax                       int
	Energy                      int
	EnergyMax                   int
	Initiative                  int
	Discernment                 int
	AP                          typ.AccountStatsStat
	MP                          typ.AccountStatsStat
	Strength                    typ.AccountStatsStat
	Vitality                    typ.AccountStatsStat
	Wisdom                      typ.AccountStatsStat
	Chance                      typ.AccountStatsStat
	Agility                     typ.AccountStatsStat
	Intelligence                typ.AccountStatsStat
	RangeModerator              typ.AccountStatsStat
	MaxSummonedCreatures        typ.AccountStatsStat
	BonusDamage                 typ.AccountStatsStat
	PhysicalBonusDamage         typ.AccountStatsStat
	WeaponSkillBonus            typ.AccountStatsStat
	BonusDamagePercent          typ.AccountStatsStat
	HealBonus                   typ.AccountStatsStat
	TrapBonus                   typ.AccountStatsStat
	TrapBonusPercent            typ.AccountStatsStat
	DamageReflected             typ.AccountStatsStat
	CriticalHitBonus            typ.AccountStatsStat
	CriticalFailureBonus        typ.AccountStatsStat
	DodgeAP                     typ.AccountStatsStat
	DodgeMP                     typ.AccountStatsStat
	NeutralResistanceFixed      typ.AccountStatsStat
	NeutralResistancePercent    typ.AccountStatsStat
	NeutralResistanceFixedPVP   typ.AccountStatsStat
	NeutralResistancePercentPVP typ.AccountStatsStat
	EarthResistanceFixed        typ.AccountStatsStat
	EarthResistancePercent      typ.AccountStatsStat
	EarthResistanceFixedPVP     typ.AccountStatsStat
	EarthResistancePercentPVP   typ.AccountStatsStat
	WaterResistanceFixed        typ.AccountStatsStat
	WaterResistancePercent      typ.AccountStatsStat
	WaterResistanceFixedPVP     typ.AccountStatsStat
	WaterResistancePercentPVP   typ.AccountStatsStat
	AirResistanceFixed          typ.AccountStatsStat
	AirResistancePercent        typ.AccountStatsStat
	AirResistanceFixedPVP       typ.AccountStatsStat
	AirResistancePercentPVP     typ.AccountStatsStat
	FireResistanceFixed         typ.AccountStatsStat
	FireResistancePercent       typ.AccountStatsStat
	FireResistanceFixedPVP      typ.AccountStatsStat
	FireResistancePercentPVP    typ.AccountStatsStat
}

func (m AccountStats) ProtocolId() d1proto.MsgSvrId {
	return d1proto.AccountStats
}

func (m AccountStats) Serialized() (string, error) {
	var alignmentEnabled int
	if m.AlignmentEnabled {
		alignmentEnabled = 1
	}

	var stats []string

	ap, err := m.AP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, ap)

	mp, err := m.MP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, mp)

	strength, err := m.Strength.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, strength)

	vitality, err := m.Vitality.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, vitality)

	wisdom, err := m.Wisdom.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, wisdom)

	chance, err := m.Chance.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, chance)

	agility, err := m.Agility.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, agility)

	intelligence, err := m.Intelligence.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, intelligence)

	rangeModerator, err := m.RangeModerator.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, rangeModerator)

	maxSummonedCreatures, err := m.MaxSummonedCreatures.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, maxSummonedCreatures)

	bonusDamage, err := m.BonusDamage.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, bonusDamage)

	physicalBonusDamage, err := m.PhysicalBonusDamage.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, physicalBonusDamage)

	weaponSkillBonus, err := m.WeaponSkillBonus.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, weaponSkillBonus)

	bonusDamagePercent, err := m.BonusDamagePercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, bonusDamagePercent)

	healBonus, err := m.HealBonus.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, healBonus)

	trapBonus, err := m.TrapBonus.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, trapBonus)

	trapBonusPercent, err := m.TrapBonusPercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, trapBonusPercent)

	damageReflected, err := m.DamageReflected.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, damageReflected)

	criticalHitBonus, err := m.CriticalHitBonus.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, criticalHitBonus)

	criticalFailureBonus, err := m.CriticalFailureBonus.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, criticalFailureBonus)

	dodgeAP, err := m.DodgeAP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, dodgeAP)

	dodgeMP, err := m.DodgeMP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, dodgeMP)

	neutralResistanceFixed, err := m.NeutralResistanceFixed.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, neutralResistanceFixed)

	neutralResistancePercent, err := m.NeutralResistancePercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, neutralResistancePercent)

	neutralResistanceFixedPVP, err := m.NeutralResistanceFixedPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, neutralResistanceFixedPVP)

	neutralResistancePercentPVP, err := m.NeutralResistancePercentPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, neutralResistancePercentPVP)

	earthResistanceFixed, err := m.EarthResistanceFixed.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, earthResistanceFixed)

	earthResistancePercent, err := m.EarthResistancePercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, earthResistancePercent)

	earthResistanceFixedPVP, err := m.EarthResistanceFixedPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, earthResistanceFixedPVP)

	earthResistancePercentPVP, err := m.EarthResistancePercentPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, earthResistancePercentPVP)

	waterResistanceFixed, err := m.WaterResistanceFixed.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, waterResistanceFixed)

	waterResistancePercent, err := m.WaterResistancePercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, waterResistancePercent)

	waterResistanceFixedPVP, err := m.WaterResistanceFixedPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, waterResistanceFixedPVP)

	waterResistancePercentPVP, err := m.WaterResistancePercentPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, waterResistancePercentPVP)

	airResistanceFixed, err := m.AirResistanceFixed.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, airResistanceFixed)

	airResistancePercent, err := m.AirResistancePercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, airResistancePercent)

	airResistanceFixedPVP, err := m.AirResistanceFixedPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, airResistanceFixedPVP)

	airResistancePercentPVP, err := m.AirResistancePercentPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, airResistancePercentPVP)

	fireResistanceFixed, err := m.FireResistanceFixed.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, fireResistanceFixed)

	fireResistancePercent, err := m.FireResistancePercent.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, fireResistancePercent)

	fireResistanceFixedPVP, err := m.FireResistanceFixedPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, fireResistanceFixedPVP)

	fireResistancePercentPVP, err := m.FireResistancePercentPVP.Serialized()
	if err != nil {
		return "", err
	}
	stats = append(stats, fireResistancePercentPVP)

	return fmt.Sprintf("%d,%d,%d|%d|%d|%d|%d~%d,%d,%d,%d,%d|%d,%d|%d,%d|%d|%d|%s",
		m.XP, m.XPLow, m.XPHigh, m.Kama, m.BonusPoints, m.BonusPointsSpell, m.Alignment, m.FakeAlignment, m.RankValue,
		m.Honour, m.Disgrace, alignmentEnabled, m.LP, m.LPMax, m.Energy, m.EnergyMax, m.Initiative, m.Discernment,
		strings.Join(stats, "|"),
	), nil
}

func (m *AccountStats) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) < 51 {
		return d1proto.ErrInvalidMsg
	}

	sli2 := strings.Split(sli[0], ",")
	if len(sli2) < 3 {
		return d1proto.ErrInvalidMsg
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
		return d1proto.ErrInvalidMsg
	}

	if sli3[0] != "" {
		sli4 := strings.Split(sli3[0], "~")
		if len(sli4) < 2 {
			return d1proto.ErrInvalidMsg
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
		rankValue, err := strconv.ParseInt(sli3[1], 10, 32)
		if err != nil {
			return err
		}
		m.RankValue = int(rankValue)
	}

	if sli3[2] != "" {
		alignmentUnknownValue, err := strconv.ParseInt(sli3[2], 10, 32)
		if err != nil {
			return err
		}
		m.AlignmentUnknownValue = int(alignmentUnknownValue)
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
		return d1proto.ErrInvalidMsg
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
		return d1proto.ErrInvalidMsg
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

	if sli[9] != "" {
		ap := &typ.AccountStatsStat{}
		err := ap.Deserialize(sli[9])
		if err != nil {
			return err
		}
		m.AP = *ap
	}

	if sli[10] != "" {
		mp := &typ.AccountStatsStat{}
		err := mp.Deserialize(sli[10])
		if err != nil {
			return err
		}
		m.MP = *mp
	}

	if sli[11] != "" {
		strength := &typ.AccountStatsStat{}
		err := strength.Deserialize(sli[11])
		if err != nil {
			return err
		}
		m.Strength = *strength
	}

	if sli[12] != "" {
		vitality := &typ.AccountStatsStat{}
		err := vitality.Deserialize(sli[12])
		if err != nil {
			return err
		}
		m.Vitality = *vitality
	}

	if sli[13] != "" {
		wisdom := &typ.AccountStatsStat{}
		err := wisdom.Deserialize(sli[13])
		if err != nil {
			return err
		}
		m.Wisdom = *wisdom
	}

	if sli[14] != "" {
		chance := &typ.AccountStatsStat{}
		err := chance.Deserialize(sli[14])
		if err != nil {
			return err
		}
		m.Chance = *chance
	}

	if sli[15] != "" {
		agility := &typ.AccountStatsStat{}
		err := agility.Deserialize(sli[15])
		if err != nil {
			return err
		}
		m.Agility = *agility
	}

	if sli[16] != "" {
		intelligence := &typ.AccountStatsStat{}
		err := intelligence.Deserialize(sli[16])
		if err != nil {
			return err
		}
		m.Intelligence = *intelligence
	}

	if sli[17] != "" {
		rangeModerator := &typ.AccountStatsStat{}
		err := rangeModerator.Deserialize(sli[17])
		if err != nil {
			return err
		}
		m.RangeModerator = *rangeModerator
	}

	if sli[18] != "" {
		maxSummonedCreatures := &typ.AccountStatsStat{}
		err := maxSummonedCreatures.Deserialize(sli[18])
		if err != nil {
			return err
		}
		m.MaxSummonedCreatures = *maxSummonedCreatures
	}

	if sli[19] != "" {
		bonusDamage := &typ.AccountStatsStat{}
		err := bonusDamage.Deserialize(sli[19])
		if err != nil {
			return err
		}
		m.BonusDamage = *bonusDamage
	}

	if sli[20] != "" {
		physicalBonusDamage := &typ.AccountStatsStat{}
		err := physicalBonusDamage.Deserialize(sli[20])
		if err != nil {
			return err
		}
		m.PhysicalBonusDamage = *physicalBonusDamage
	}

	if sli[21] != "" {
		weaponSkillBonus := &typ.AccountStatsStat{}
		err := weaponSkillBonus.Deserialize(sli[21])
		if err != nil {
			return err
		}
		m.WeaponSkillBonus = *weaponSkillBonus
	}

	if sli[22] != "" {
		bonusDamagePercent := &typ.AccountStatsStat{}
		err := bonusDamagePercent.Deserialize(sli[22])
		if err != nil {
			return err
		}
		m.BonusDamagePercent = *bonusDamagePercent
	}

	if sli[23] != "" {
		healBonus := &typ.AccountStatsStat{}
		err := healBonus.Deserialize(sli[23])
		if err != nil {
			return err
		}
		m.HealBonus = *healBonus
	}

	if sli[24] != "" {
		trapBonus := &typ.AccountStatsStat{}
		err := trapBonus.Deserialize(sli[24])
		if err != nil {
			return err
		}
		m.TrapBonus = *trapBonus
	}

	if sli[25] != "" {
		trapBonusPercent := &typ.AccountStatsStat{}
		err := trapBonusPercent.Deserialize(sli[25])
		if err != nil {
			return err
		}
		m.TrapBonusPercent = *trapBonusPercent
	}

	if sli[26] != "" {
		damageReflected := &typ.AccountStatsStat{}
		err := damageReflected.Deserialize(sli[26])
		if err != nil {
			return err
		}
		m.DamageReflected = *damageReflected
	}

	if sli[27] != "" {
		criticalHitBonus := &typ.AccountStatsStat{}
		err := criticalHitBonus.Deserialize(sli[27])
		if err != nil {
			return err
		}
		m.CriticalHitBonus = *criticalHitBonus
	}

	if sli[28] != "" {
		criticalFailureBonus := &typ.AccountStatsStat{}
		err := criticalFailureBonus.Deserialize(sli[28])
		if err != nil {
			return err
		}
		m.CriticalFailureBonus = *criticalFailureBonus
	}

	if sli[29] != "" {
		dodgeAP := &typ.AccountStatsStat{}
		err := dodgeAP.Deserialize(sli[29])
		if err != nil {
			return err
		}
		m.DodgeAP = *dodgeAP
	}

	if sli[30] != "" {
		dodgeMP := &typ.AccountStatsStat{}
		err := dodgeMP.Deserialize(sli[30])
		if err != nil {
			return err
		}
		m.DodgeMP = *dodgeMP
	}

	if sli[31] != "" {
		neutralResistanceFixed := &typ.AccountStatsStat{}
		err := neutralResistanceFixed.Deserialize(sli[31])
		if err != nil {
			return err
		}
		m.NeutralResistanceFixed = *neutralResistanceFixed
	}

	if sli[32] != "" {
		neutralResistancePercent := &typ.AccountStatsStat{}
		err := neutralResistancePercent.Deserialize(sli[32])
		if err != nil {
			return err
		}
		m.NeutralResistancePercent = *neutralResistancePercent
	}

	if sli[33] != "" {
		neutralResistanceFixedPVP := &typ.AccountStatsStat{}
		err := neutralResistanceFixedPVP.Deserialize(sli[33])
		if err != nil {
			return err
		}
		m.NeutralResistanceFixedPVP = *neutralResistanceFixedPVP
	}

	if sli[34] != "" {
		neutralResistancePercentPVP := &typ.AccountStatsStat{}
		err := neutralResistancePercentPVP.Deserialize(sli[34])
		if err != nil {
			return err
		}
		m.NeutralResistancePercentPVP = *neutralResistancePercentPVP
	}

	if sli[35] != "" {
		earthResistanceFixed := &typ.AccountStatsStat{}
		err := earthResistanceFixed.Deserialize(sli[35])
		if err != nil {
			return err
		}
		m.EarthResistanceFixed = *earthResistanceFixed
	}

	if sli[36] != "" {
		earthResistancePercent := &typ.AccountStatsStat{}
		err := earthResistancePercent.Deserialize(sli[36])
		if err != nil {
			return err
		}
		m.EarthResistancePercent = *earthResistancePercent
	}

	if sli[37] != "" {
		earthResistanceFixedPVP := &typ.AccountStatsStat{}
		err := earthResistanceFixedPVP.Deserialize(sli[37])
		if err != nil {
			return err
		}
		m.EarthResistanceFixedPVP = *earthResistanceFixedPVP
	}

	if sli[38] != "" {
		earthResistancePercentPVP := &typ.AccountStatsStat{}
		err := earthResistancePercentPVP.Deserialize(sli[38])
		if err != nil {
			return err
		}
		m.EarthResistancePercentPVP = *earthResistancePercentPVP
	}

	if sli[39] != "" {
		waterResistanceFixed := &typ.AccountStatsStat{}
		err := waterResistanceFixed.Deserialize(sli[39])
		if err != nil {
			return err
		}
		m.WaterResistanceFixed = *waterResistanceFixed
	}

	if sli[40] != "" {
		waterResistancePercent := &typ.AccountStatsStat{}
		err := waterResistancePercent.Deserialize(sli[40])
		if err != nil {
			return err
		}
		m.WaterResistancePercent = *waterResistancePercent
	}

	if sli[41] != "" {
		waterResistanceFixedPVP := &typ.AccountStatsStat{}
		err := waterResistanceFixedPVP.Deserialize(sli[41])
		if err != nil {
			return err
		}
		m.WaterResistanceFixedPVP = *waterResistanceFixedPVP
	}

	if sli[42] != "" {
		waterResistancePercentPVP := &typ.AccountStatsStat{}
		err := waterResistancePercentPVP.Deserialize(sli[42])
		if err != nil {
			return err
		}
		m.WaterResistancePercentPVP = *waterResistancePercentPVP
	}

	if sli[43] != "" {
		airResistanceFixed := &typ.AccountStatsStat{}
		err := airResistanceFixed.Deserialize(sli[43])
		if err != nil {
			return err
		}
		m.AirResistanceFixed = *airResistanceFixed
	}

	if sli[44] != "" {
		airResistancePercent := &typ.AccountStatsStat{}
		err := airResistancePercent.Deserialize(sli[44])
		if err != nil {
			return err
		}
		m.AirResistancePercent = *airResistancePercent
	}

	if sli[45] != "" {
		airResistanceFixedPVP := &typ.AccountStatsStat{}
		err := airResistanceFixedPVP.Deserialize(sli[45])
		if err != nil {
			return err
		}
		m.AirResistanceFixedPVP = *airResistanceFixedPVP
	}

	if sli[46] != "" {
		airResistancePercentPVP := &typ.AccountStatsStat{}
		err := airResistancePercentPVP.Deserialize(sli[46])
		if err != nil {
			return err
		}
		m.AirResistancePercentPVP = *airResistancePercentPVP
	}

	if sli[47] != "" {
		fireResistanceFixed := &typ.AccountStatsStat{}
		err := fireResistanceFixed.Deserialize(sli[47])
		if err != nil {
			return err
		}
		m.FireResistanceFixed = *fireResistanceFixed
	}

	if sli[48] != "" {
		fireResistancePercent := &typ.AccountStatsStat{}
		err := fireResistancePercent.Deserialize(sli[48])
		if err != nil {
			return err
		}
		m.FireResistancePercent = *fireResistancePercent
	}

	if sli[49] != "" {
		fireResistanceFixedPVP := &typ.AccountStatsStat{}
		err := fireResistanceFixedPVP.Deserialize(sli[49])
		if err != nil {
			return err
		}
		m.FireResistanceFixedPVP = *fireResistanceFixedPVP
	}

	if sli[50] != "" {
		fireResistancePercentPVP := &typ.AccountStatsStat{}
		err := fireResistancePercentPVP.Deserialize(sli[50])
		if err != nil {
			return err
		}
		m.FireResistancePercentPVP = *fireResistancePercentPVP
	}

	return nil
}
