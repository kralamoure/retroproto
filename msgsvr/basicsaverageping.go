// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAveragePing struct{}

func (m BasicsAveragePing) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.BasicsAveragePing
}

func (m BasicsAveragePing) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAveragePing) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
