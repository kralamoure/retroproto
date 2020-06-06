package enum

var GameMovementSpriteType = newGameMovementSpriteType()

func newGameMovementSpriteType() *gameMovementSpriteType {
	return &gameMovementSpriteType{
		Creature:         -1,
		Monster:          -2,
		MonsterGroup:     -3,
		NPC:              -4,
		OfflineCharacter: -5,
		TaxCollector:     -6,
		Mutant:           -7,
		MutantPlayer:     -8,
		ParkMount:        -9,
		Prism:            -10,
	}
}

type gameMovementSpriteType struct {
	Creature         int
	Monster          int
	MonsterGroup     int
	NPC              int
	OfflineCharacter int
	TaxCollector     int
	Mutant           int
	MutantPlayer     int
	ParkMount        int
	Prism            int
}
