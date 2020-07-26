package enum

var ItemsAdd = newItemsAdd()

func newItemsAdd() *itemsAdd {
	return &itemsAdd{
		InventoryFull:      'F',
		TooLowLevelForItem: 'L',
		AlreadyEquipped:    'A',
	}
}

type itemsAdd struct {
	InventoryFull      rune
	TooLowLevelForItem rune
	AlreadyEquipped    rune
}
