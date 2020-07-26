package enum

var ItemsAddSuccessItemType = newItemsAddSuccessItemType()

func newItemsAddSuccessItemType() *itemsAddSuccessItemType {
	return &itemsAddSuccessItemType{
		G:       'G',
		Objects: 'O',
	}
}

type itemsAddSuccessItemType struct {
	// I don't know what 'G' stands for. Maybe Gift, Game, General, etc.
	G       rune
	Objects rune
}
