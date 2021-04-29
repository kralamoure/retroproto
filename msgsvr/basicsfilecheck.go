// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsFileCheck struct{}

func (m BasicsFileCheck) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsFileCheck
}

func (m BasicsFileCheck) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsFileCheck) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
