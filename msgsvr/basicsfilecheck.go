// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsFileCheck struct{}

func (m BasicsFileCheck) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsFileCheck
}

func (m BasicsFileCheck) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsFileCheck) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
