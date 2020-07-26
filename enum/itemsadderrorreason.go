package enum

var ItemsAddErrorReason = newItemsAddErrorReason()

func newItemsAddErrorReason() *itemsAddErrorReason {
	return &itemsAddErrorReason{
		InventoryFull:      'F',
		TooLowLevelForItem: 'L',
		AlreadyEquipped:    'A',
	}
}

type itemsAddErrorReason struct {
	InventoryFull      rune
	TooLowLevelForItem rune
	AlreadyEquipped    rune
}
