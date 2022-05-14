package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type BasicsSubscriberRestrictionRemove struct{}

func (m BasicsSubscriberRestrictionRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsSubscriberRestrictionRemove
}

func (m BasicsSubscriberRestrictionRemove) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsSubscriberRestrictionRemove) Deserialize(extra string) error {
	return nil
}
