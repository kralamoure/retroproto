package typ

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type CommonRestrictions struct {
	CantAssault                          bool
	CantChallenge                        bool
	CantExchange                         bool
	CantAttack                           bool
	CantChatToAll                        bool
	CantBeMerchant                       bool
	CantUseObject                        bool
	CantInteractWithTaxCollector         bool
	CantUseInteractiveObjects            bool
	CantSpeakNPC                         bool
	CantAttackDungeonMonstersWhenMutant  bool
	CantMoveInAllDirections              bool
	CantAttackMonstersAnywhereWhenMutant bool
	CantInteractWithPrism                bool
}

func NewCommonRestrictions(extra string) (CommonRestrictions, error) {
	var m CommonRestrictions

	if err := m.Deserialize(extra); err != nil {
		return CommonRestrictions{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m CommonRestrictions) Serialized() (string, error) {
	var restrictions int

	if m.CantAssault {
		restrictions += 1 << 0
	}

	if m.CantChallenge {
		restrictions += 1 << 1
	}

	if m.CantExchange {
		restrictions += 1 << 2
	}

	if !m.CantAttack {
		restrictions += 1 << 3
	}

	if m.CantChatToAll {
		restrictions += 1 << 4
	}

	if m.CantBeMerchant {
		restrictions += 1 << 5
	}

	if m.CantUseObject {
		restrictions += 1 << 6
	}

	if m.CantInteractWithTaxCollector {
		restrictions += 1 << 7
	}

	if m.CantUseInteractiveObjects {
		restrictions += 1 << 8
	}

	if m.CantSpeakNPC {
		restrictions += 1 << 9
	}

	if !m.CantAttackDungeonMonstersWhenMutant {
		restrictions += 1 << 12
	}

	if !m.CantMoveInAllDirections {
		restrictions += 1 << 13
	}

	if !m.CantAttackMonstersAnywhereWhenMutant {
		restrictions += 1 << 14
	}

	if m.CantInteractWithPrism {
		restrictions += 1 << 15
	}

	return strconv.FormatInt(int64(restrictions), 36), nil
}

func (m *CommonRestrictions) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	restrictionsN, err := strconv.ParseInt(extra, 36, 32)
	if err != nil {
		return err
	}
	restrictions := int(restrictionsN)

	canAssault := int(1 << 0)
	m.CantAssault = !(restrictions&canAssault != canAssault)

	canChallenge := int(1 << 1)
	m.CantChallenge = !(restrictions&canChallenge != canChallenge)

	canExchange := int(1 << 2)
	m.CantExchange = !(restrictions&canExchange != canExchange)

	canAttack := int(1 << 3)
	m.CantAttack = !(restrictions&canAttack == canAttack)

	canChatToAll := int(1 << 4)
	m.CantChatToAll = !(restrictions&canChatToAll != canChatToAll)

	canBeMerchant := int(1 << 5)
	m.CantBeMerchant = !(restrictions&canBeMerchant != canBeMerchant)

	canUseObject := int(1 << 6)
	m.CantUseObject = !(restrictions&canUseObject != canUseObject)

	cantInteractWithTaxCollector := int(1 << 7)
	m.CantInteractWithTaxCollector = restrictions&cantInteractWithTaxCollector == cantInteractWithTaxCollector

	canUseInteractiveObjects := int(1 << 8)
	m.CantUseInteractiveObjects = !(restrictions&canUseInteractiveObjects != canUseInteractiveObjects)

	cantSpeakNPC := int(1 << 9)
	m.CantSpeakNPC = restrictions&cantSpeakNPC == cantSpeakNPC

	canAttackDungeonMonstersWhenMutant := int(1 << 12)
	m.CantAttackDungeonMonstersWhenMutant = !(restrictions&canAttackDungeonMonstersWhenMutant ==
		canAttackDungeonMonstersWhenMutant)

	canMoveInAllDirections := int(1 << 13)
	m.CantMoveInAllDirections = !(restrictions&(canMoveInAllDirections) == canMoveInAllDirections)

	canAttackMonstersAnywhereWhenMutant := int(1 << 14)
	m.CantAttackMonstersAnywhereWhenMutant = !(restrictions&canAttackMonstersAnywhereWhenMutant ==
		canAttackMonstersAnywhereWhenMutant)

	cantInteractWithPrism := int(1 << 15)
	m.CantInteractWithPrism = restrictions&cantInteractWithPrism == cantInteractWithPrism

	return nil
}
