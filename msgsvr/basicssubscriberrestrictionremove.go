package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsSubscriberRestrictionRemove struct{}

func NewBasicsSubscriberRestrictionRemove(extra string) (BasicsSubscriberRestrictionRemove, error) {
	var m BasicsSubscriberRestrictionRemove

	if err := m.Deserialize(extra); err != nil {
		return BasicsSubscriberRestrictionRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsSubscriberRestrictionRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsSubscriberRestrictionRemove
}

func (m BasicsSubscriberRestrictionRemove) MessageName() string {
	return "BasicsSubscriberRestrictionRemove"
}

func (m BasicsSubscriberRestrictionRemove) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsSubscriberRestrictionRemove) Deserialize(extra string) error {
	return nil
}
