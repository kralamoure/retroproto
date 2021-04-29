// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAveragePing struct{}

func (m BasicsAveragePing) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsAveragePing
}

func (m BasicsAveragePing) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsAveragePing) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
