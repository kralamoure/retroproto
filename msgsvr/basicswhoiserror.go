// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsWhoIsError struct{}

func (m BasicsWhoIsError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsWhoIsError
}

func (m BasicsWhoIsError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsWhoIsError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
