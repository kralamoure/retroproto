package enum

var MountEquipErrorReason = newMountEquipErrorReason()

func newMountEquipErrorReason() *mountEquipErrorReason {
	return &mountEquipErrorReason{
		InventoryNotEmpty: '-',
		AlreadyHaveOne:    '+',
		Ride:              'r',
	}
}

type mountEquipErrorReason struct {
	InventoryNotEmpty rune
	AlreadyHaveOne    rune
	Ride              rune
}
