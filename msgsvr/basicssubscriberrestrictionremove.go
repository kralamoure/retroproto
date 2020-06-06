package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsSubscriberRestrictionRemove struct{}

func (m BasicsSubscriberRestrictionRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsSubscriberRestrictionRemove
}

func (m BasicsSubscriberRestrictionRemove) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsSubscriberRestrictionRemove) Deserialize(extra string) error {
	return nil
}
