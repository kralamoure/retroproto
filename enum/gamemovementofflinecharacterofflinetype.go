package enum

var GameMovementOfflineCharacterOfflineType = newGameMovementOfflineCharacterOfflineType()

func newGameMovementOfflineCharacterOfflineType() *gameMovementOfflineCharacterOfflineType {
	return &gameMovementOfflineCharacterOfflineType{
		ShopAll:       0,
		ShopEquipment: 1,
		ShopMisc:      2,
		ShopResources: 3,
	}
}

type gameMovementOfflineCharacterOfflineType struct {
	ShopAll       int
	ShopEquipment int
	ShopMisc      int
	ShopResources int
}
