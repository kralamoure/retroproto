package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsSubscriberRestrictionRemove struct{}

func (m BasicsSubscriberRestrictionRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsSubscriberRestrictionRemove
}

func (m BasicsSubscriberRestrictionRemove) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsSubscriberRestrictionRemove) Deserialize(extra string) error {
	return nil
}
