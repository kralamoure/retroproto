package enum

var BasicsSubscriberRestrictionAddDialogId = newBasicsSubscriberRestrictionAddDialogId()

func newBasicsSubscriberRestrictionAddDialogId() *basicsSubscriberRestrictionAddDialogId {
	return &basicsSubscriberRestrictionAddDialogId{
		Default: 10,
	}
}

type basicsSubscriberRestrictionAddDialogId struct {
	Default int
}
